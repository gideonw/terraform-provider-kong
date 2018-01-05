package kong

// import (
// 	"fmt"
// 	"github.com/hashicorp/terraform/helper/resource"
// 	"github.com/hashicorp/terraform/terraform"
// 	"github.com/gideonw/gokong"
// 	"testing"
// )

// func TestAccKongConsumerJWT(t *testing.T) {

// 	resource.Test(t, resource.TestCase{
// 		Providers:    testAccProviders,
// 		CheckDestroy: testAccCheckKongConsumerJWTDestroy,
// 		Steps: []resource.TestStep{
// 			{
// 				Config: testCreateConsumerConfig,
// 				Check: resource.ComposeTestCheckFunc(
// 					testAccCheckKongConsumerJWTExists("kong_consumer.consumer"),
// 					resource.TestCheckResourceAttr("kong_consumer.consumer", "username", "User1"),
// 					resource.TestCheckResourceAttr("kong_consumer.consumer", "custom_id", "123"),
// 				),
// 			},
// 			{
// 				Config: testUpdateConsumerConfig,
// 				Check: resource.ComposeTestCheckFunc(
// 					testAccCheckKongConsumerJWTExists("kong_consumer.consumer"),
// 					resource.TestCheckResourceAttr("kong_consumer.consumer", "username", "User2"),
// 					resource.TestCheckResourceAttr("kong_consumer.consumer", "custom_id", "456"),
// 				),
// 			},
// 		},
// 	})
// }

// func testAccCheckKongConsumerJWTDestroy(state *terraform.State) error {

// 	client := testAccProvider.Meta().(*gokong.KongAdminClient)

// 	consumers := getResourcesByType("kong_consumer", state)

// 	if len(consumers) != 1 {
// 		return fmt.Errorf("expecting only 1 consumer resource found %v", len(consumers))
// 	}

// 	response, err := client.Consumers().GetByID(consumers[0].Primary.ID)

// 	if err != nil {
// 		return fmt.Errorf("error calling get consumer by id: %v", err)
// 	}

// 	if response != nil {
// 		return fmt.Errorf("consumer %s still exists, %+v", consumers[0].Primary.ID, response)
// 	}

// 	return nil
// }

// func testAccCheckKongConsumerJWTExists(resourceKey string) resource.TestCheckFunc {

// 	return func(s *terraform.State) error {
// 		rs, ok := s.RootModule().Resources[resourceKey]

// 		if !ok {
// 			return fmt.Errorf("not found: %s", resourceKey)
// 		}

// 		if rs.Primary.ID == "" {
// 			return fmt.Errorf("no ID is set")
// 		}

// 		client := testAccProvider.Meta().(*gokong.KongAdminClient)

// 		api, err := client.Consumers().GetByID(rs.Primary.ID)

// 		if err != nil {
// 			return err
// 		}

// 		if api == nil {
// 			return fmt.Errorf("consumer with id %v not found", rs.Primary.ID)
// 		}

// 		return nil
// 	}
// }

// const testCreateConsumerConfig = `
// resource "kong_consumer" "consumer" {
// 	username  = "User1"
// 	custom_id = "123"
// }
// `
// const testUpdateConsumerConfig = `
// resource "kong_consumer" "consumer" {
// 	username  = "User2"
// 	custom_id = "456"
// }
// `

// func TestAccKongConsumerJWT(t *testing.T) {

// 	resource.Test(t, resource.TestCase{
// 		Providers:    testAccProviders,
// 		CheckDestroy: testAccCheckKongConsumerJWTDestroy,
// 		Steps: []resource.TestStep{
// 			{
// 				Config: testCreateConsumerConfig,
// 				Check: resource.ComposeTestCheckFunc(
// 					testAccCheckKongConsumerJWTExists("kong_consumer.consumer"),
// 					resource.TestCheckResourceAttr("kong_consumer.consumer", "username", "User1"),
// 					resource.TestCheckResourceAttr("kong_consumer.consumer", "custom_id", "123"),
// 				),
// 			},
// 			{
// 				Config: testUpdateConsumerConfig,
// 				Check: resource.ComposeTestCheckFunc(
// 					testAccCheckKongConsumerJWTExists("kong_consumer.consumer"),
// 					resource.TestCheckResourceAttr("kong_consumer.consumer", "username", "User2"),
// 					resource.TestCheckResourceAttr("kong_consumer.consumer", "custom_id", "456"),
// 				),
// 			},
// 		},
// 	})
// }

// func testAccCheckKongConsumerJWTDestroy(state *terraform.State) error {

// 	client := testAccProvider.Meta().(*gokong.KongAdminClient)

// 	consumers := getResourcesByType("kong_consumer", state)

// 	if len(consumers) != 1 {
// 		return fmt.Errorf("expecting only 1 consumer resource found %v", len(consumers))
// 	}

// 	response, err := client.Consumers().GetByID(consumers[0].Primary.ID)

// 	if err != nil {
// 		return fmt.Errorf("error calling get consumer by id: %v", err)
// 	}

// 	if response != nil {
// 		return fmt.Errorf("consumer %s still exists, %+v", consumers[0].Primary.ID, response)
// 	}

// 	return nil
// }

// func testAccCheckKongConsumerJWTExists(resourceKey string) resource.TestCheckFunc {

// 	return func(s *terraform.State) error {
// 		rs, ok := s.RootModule().Resources[resourceKey]

// 		if !ok {
// 			return fmt.Errorf("not found: %s", resourceKey)
// 		}

// 		if rs.Primary.ID == "" {
// 			return fmt.Errorf("no ID is set")
// 		}

// 		client := testAccProvider.Meta().(*gokong.KongAdminClient)

// 		api, err := client.Consumers().GetByID(rs.Primary.ID)

// 		if err != nil {
// 			return err
// 		}

// 		if api == nil {
// 			return fmt.Errorf("consumer with id %v not found", rs.Primary.ID)
// 		}

// 		return nil
// 	}
// }

// const testCreateConsumerConfig = `
// resource "kong_consumer" "consumer" {
// 	username  = "User1"
// 	custom_id = "123"
// }
// `
// const testUpdateConsumerConfig = `
// resource "kong_consumer" "consumer" {
// 	username  = "User2"
// 	custom_id = "456"
// }
// `
