package kong

import (
	"fmt"

	"github.com/gideonw/gokong"
	"github.com/hashicorp/terraform/helper/schema"
)

func resourceKongConsumer() *schema.Resource {
	return &schema.Resource{
		Create: resourceKongConsumerCreate,
		Read:   resourceKongConsumerRead,
		Delete: resourceKongConsumerDelete,
		Update: resourceKongConsumerUpdate,

		Schema: map[string]*schema.Schema{
			"username": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: false,
			},
			"custom_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: false,
			},
		},
	}
}

func resourceKongConsumerCreate(d *schema.ResourceData, meta interface{}) error {

	consumerRequest := createKongConsumerRequestFromResourceData(d)

	consumer, err := meta.(*gokong.KongAdminClient).Consumers().CreateOrUpdate(consumerRequest)

	if err != nil {
		return fmt.Errorf("failed to create kong consumer: %v error: %v", consumerRequest, err)
	}

	d.SetId(consumer.ID)

	return resourceKongConsumerRead(d, meta)
}

func resourceKongConsumerUpdate(d *schema.ResourceData, meta interface{}) error {
	d.Partial(false)

	consumerRequest := createKongConsumerRequestFromResourceData(d)

	_, err := meta.(*gokong.KongAdminClient).Consumers().UpdateByID(d.Id(), consumerRequest)

	if err != nil {
		return fmt.Errorf("error updating kong consumer: %s", err)
	}

	return resourceKongConsumerRead(d, meta)
}

func resourceKongConsumerRead(d *schema.ResourceData, meta interface{}) error {

	id := d.Id()
	consumer, err := meta.(*gokong.KongAdminClient).Consumers().GetByID(id)

	if err != nil {
		return fmt.Errorf("could not find kong consumer with id: %s error: %v", id, err)
	}

	d.Set("username", consumer.Username)
	d.Set("custom_id", consumer.CustomID)

	return nil
}

func resourceKongConsumerDelete(d *schema.ResourceData, meta interface{}) error {

	err := meta.(*gokong.KongAdminClient).Consumers().DeleteByID(d.Id())

	if err != nil {
		return fmt.Errorf("could not delete kong consumer: %v", err)
	}

	return nil
}

func createKongConsumerRequestFromResourceData(d *schema.ResourceData) *gokong.ConsumerRequest {

	consumerRequest := &gokong.ConsumerRequest{}

	consumerRequest.Username = readStringFromResource(d, "username")
	consumerRequest.CustomID = readStringFromResource(d, "custom_id")

	return consumerRequest
}
