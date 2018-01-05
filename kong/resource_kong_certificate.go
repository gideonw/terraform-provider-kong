package kong

import (
	"fmt"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/gideonw/gokong"
)

func resourceKongCertificate() *schema.Resource {
	return &schema.Resource{
		Create: resourceKongCertificateCreate,
		Read:   resourceKongCertificateRead,
		Delete: resourceKongCertificateDelete,
		Update: resourceKongCertificateUpdate,

		Schema: map[string]*schema.Schema{
			"certificate": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: false,
			},
			"private_key": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: false,
			},
		},
	}
}

func resourceKongCertificateCreate(d *schema.ResourceData, meta interface{}) error {

	certificateRequest := createKongCertificateRequestFromResourceData(d)

	certificate, err := meta.(*gokong.KongAdminClient).Certificates().Create(certificateRequest)

	if err != nil {
		return fmt.Errorf("failed to create kong certificate: %v error: %v", certificateRequest, err)
	}

	d.SetId(certificate.ID)

	return resourceKongCertificateRead(d, meta)
}

func resourceKongCertificateUpdate(d *schema.ResourceData, meta interface{}) error {
	d.Partial(false)

	certificateRequest := createKongCertificateRequestFromResourceData(d)

	_, err := meta.(*gokong.KongAdminClient).Certificates().UpdateByID(d.Id(), certificateRequest)

	if err != nil {
		return fmt.Errorf("error updating kong certificate: %s", err)
	}

	return resourceKongCertificateRead(d, meta)
}

func resourceKongCertificateRead(d *schema.ResourceData, meta interface{}) error {

	certificate, err := meta.(*gokong.KongAdminClient).Certificates().GetByID(d.Id())

	if err != nil {
		return fmt.Errorf("could not find kong certificate: %v", err)
	}

	d.Set("certificate", certificate.Cert)
	d.Set("private_key", certificate.Key)

	return nil
}

func resourceKongCertificateDelete(d *schema.ResourceData, meta interface{}) error {

	err := meta.(*gokong.KongAdminClient).Certificates().DeleteByID(d.Id())

	if err != nil {
		return fmt.Errorf("could not delete kong certificate: %v", err)
	}

	return nil
}

func createKongCertificateRequestFromResourceData(d *schema.ResourceData) *gokong.CertificateRequest {

	certificateRequest := &gokong.CertificateRequest{}

	certificateRequest.Cert = readStringFromResource(d, "certificate")
	certificateRequest.Key = readStringFromResource(d, "private_key")

	return certificateRequest
}
