package kong

import (
	"fmt"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/gideonw/gokong"
)

func resourceKongConsumerCredentialJWT() *schema.Resource {
	return &schema.Resource{
		Create: resourceKongConsumerJWTCreate,
		Read:   resourceKongConsumerJWTRead,
		Delete: resourceKongConsumerJWTDelete,

		Schema: map[string]*schema.Schema{
			"consumer_id": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"key": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"algorithm": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"rsa_public_key": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"secret": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
		},
	}
}

func resourceKongConsumerJWTCreate(d *schema.ResourceData, meta interface{}) error {
	credRequest := createKongConsumerJWTCredentialRequestFromResourceData(d)

	cred, err := meta.(*gokong.KongAdminClient).ConsumerCredentials().CreateJWTCredential(credRequest)
	if err != nil {
		return fmt.Errorf("failed to create kong consumer JWT credentials: %v error: %v", credRequest, err)
	}

	d.SetId(cred.ID)

	return resourceKongConsumerJWTRead(d, meta)
}

func resourceKongConsumerJWTRead(d *schema.ResourceData, meta interface{}) error {
	id := d.Id()
	consumerID := readStringFromResource(d, "consumer_id")

	cred, err := meta.(*gokong.KongAdminClient).ConsumerCredentials().GetJWTByID(consumerID, id)
	if err != nil {
		return fmt.Errorf("could not find kong consumer JWT credential with id: %s error: %v", id, err)
	}

	if cred == nil {
		return nil
	}

	d.Set("key", cred.Key)
	d.Set("consumer_id", cred.ConsumerID)
	d.Set("algorithm", cred.Algorithm)
	d.Set("rsa_public_key", cred.RSAPublicKey)
	d.Set("secret", cred.Secret)

	return nil
}

func resourceKongConsumerJWTDelete(d *schema.ResourceData, meta interface{}) error {
	id := d.Id()
	consumerID := readStringFromResource(d, "consumer_id")

	err := meta.(*gokong.KongAdminClient).ConsumerCredentials().DeleteJWTByID(consumerID, id)
	if err != nil {
		return fmt.Errorf("could not delete kong consumer: %v", err)
	}

	return nil
}

func createKongConsumerJWTCredentialRequestFromResourceData(d *schema.ResourceData) *gokong.ConsumerJWTCredentialRequest {
	consumerJWTCredReq := &gokong.ConsumerJWTCredentialRequest{}

	consumerJWTCredReq.Key = readStringFromResource(d, "key")
	consumerJWTCredReq.ConsumerID = readStringFromResource(d, "consumer_id")
	consumerJWTCredReq.Algorithm = readStringFromResource(d, "algorithm")
	consumerJWTCredReq.RSAPublicKey = readStringFromResource(d, "rsa_public_key")
	consumerJWTCredReq.Secret = readStringFromResource(d, "secret")

	return consumerJWTCredReq
}
