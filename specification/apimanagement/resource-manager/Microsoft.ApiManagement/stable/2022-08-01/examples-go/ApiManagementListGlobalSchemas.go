package armapimanagement_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/apimanagement/armapimanagement/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/4cd95123fb961c68740565a1efcaa5e43bd35802/specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2022-08-01/examples/ApiManagementListGlobalSchemas.json
func ExampleGlobalSchemaClient_NewListByServicePager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armapimanagement.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewGlobalSchemaClient().NewListByServicePager("rg1", "apimService1", &armapimanagement.GlobalSchemaClientListByServiceOptions{Filter: nil,
		Top:  nil,
		Skip: nil,
	})
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range page.Value {
			// You could use page here. We use blank identifier for just demo purposes.
			_ = v
		}
		// If the HTTP response code is 200 as defined in example definition, your page structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
		// page.GlobalSchemaCollection = armapimanagement.GlobalSchemaCollection{
		// 	Count: to.Ptr[int64](2),
		// 	Value: []*armapimanagement.GlobalSchemaContract{
		// 		{
		// 			Name: to.Ptr("schema1"),
		// 			Type: to.Ptr("Microsoft.ApiManagement/service/schemas"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.ApiManagement/service/apimService1/schemas/schema1"),
		// 			Properties: &armapimanagement.GlobalSchemaContractProperties{
		// 				Description: to.Ptr("sample schema description"),
		// 				SchemaType: to.Ptr(armapimanagement.SchemaTypeXML),
		// 				Value: "<xsd:schema xmlns:xsd=\"http://www.w3.org/2001/XMLSchema\"\r\n           xmlns:tns=\"http://tempuri.org/PurchaseOrderSchema.xsd\"\r\n           targetNamespace=\"http://tempuri.org/PurchaseOrderSchema.xsd\"\r\n           elementFormDefault=\"qualified\">\r\n <xsd:element name=\"PurchaseOrder\" type=\"tns:PurchaseOrderType\"/>\r\n <xsd:complexType name=\"PurchaseOrderType\">\r\n  <xsd:sequence>\r\n   <xsd:element name=\"ShipTo\" type=\"tns:USAddress\" maxOccurs=\"2\"/>\r\n   <xsd:element name=\"BillTo\" type=\"tns:USAddress\"/>\r\n  </xsd:sequence>\r\n  <xsd:attribute name=\"OrderDate\" type=\"xsd:date\"/>\r\n </xsd:complexType>\r\n\r\n <xsd:complexType name=\"USAddress\">\r\n  <xsd:sequence>\r\n   <xsd:element name=\"name\"   type=\"xsd:string\"/>\r\n   <xsd:element name=\"street\" type=\"xsd:string\"/>\r\n   <xsd:element name=\"city\"   type=\"xsd:string\"/>\r\n   <xsd:element name=\"state\"  type=\"xsd:string\"/>\r\n   <xsd:element name=\"zip\"    type=\"xsd:integer\"/>\r\n  </xsd:sequence>\r\n  <xsd:attribute name=\"country\" type=\"xsd:NMTOKEN\" fixed=\"US\"/>\r\n </xsd:complexType>\r\n</xsd:schema>",
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("schema2"),
		// 			Type: to.Ptr("Microsoft.ApiManagement/service/schemas"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.ApiManagement/service/apimService1/schemas/schema2"),
		// 			Properties: &armapimanagement.GlobalSchemaContractProperties{
		// 				Description: to.Ptr("sample schema description"),
		// 				Document: map[string]any{
		// 					"type": "object",
		// 					"$id": "https://example.com/person.schema.json",
		// 					"$schema": "https://json-schema.org/draft/2020-12/schema",
		// 					"properties":map[string]any{
		// 						"age":map[string]any{
		// 							"type": "integer",
		// 							"description": "Age in years which must be equal to or greater than zero.",
		// 							"minimum": float64(0),
		// 						},
		// 						"firstName":map[string]any{
		// 							"type": "string",
		// 							"description": "The person's first name.",
		// 						},
		// 						"lastName":map[string]any{
		// 							"type": "string",
		// 							"description": "The person's last name.",
		// 						},
		// 					},
		// 					"title": "Person",
		// 				},
		// 				SchemaType: to.Ptr(armapimanagement.SchemaTypeJSON),
		// 			},
		// 	}},
		// }
	}
}
