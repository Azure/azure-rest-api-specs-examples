package armlogic_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/logic/armlogic"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/logic/resource-manager/Microsoft.Logic/stable/2019-05-01/examples/IntegrationAccountSchemas_CreateOrUpdate.json
func ExampleIntegrationAccountSchemasClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armlogic.NewIntegrationAccountSchemasClient("34adfa4f-cedf-4dc0-ba29-b6d1a69ab345", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.CreateOrUpdate(ctx,
		"testResourceGroup",
		"testIntegrationAccount",
		"testSchema",
		armlogic.IntegrationAccountSchema{
			Location: to.Ptr("westus"),
			Tags: map[string]*string{
				"integrationAccountSchemaName": to.Ptr("IntegrationAccountSchema8120"),
			},
			Properties: &armlogic.IntegrationAccountSchemaProperties{
				Content:     to.Ptr("<?xml version=\"1.0\" encoding=\"utf-16\"?>\r\n<xs:schema xmlns:b=\"http://schemas.microsoft.com/BizTalk/2003\" xmlns=\"http://Inbound_EDI.OrderFile\" targetNamespace=\"http://Inbound_EDI.OrderFile\" xmlns:xs=\"http://www.w3.org/2001/XMLSchema\">\r\n  <xs:annotation>\r\n    <xs:appinfo>\r\n      <b:schemaInfo default_pad_char=\" \" count_positions_by_byte=\"false\" parser_optimization=\"speed\" lookahead_depth=\"3\" suppress_empty_nodes=\"false\" generate_empty_nodes=\"true\" allow_early_termination=\"false\" early_terminate_optional_fields=\"false\" allow_message_breakup_of_infix_root=\"false\" compile_parse_tables=\"false\" standard=\"Flat File\" root_reference=\"OrderFile\" />\r\n      <schemaEditorExtension:schemaInfo namespaceAlias=\"b\" extensionClass=\"Microsoft.BizTalk.FlatFileExtension.FlatFileExtension\" standardName=\"Flat File\" xmlns:schemaEditorExtension=\"http://schemas.microsoft.com/BizTalk/2003/SchemaEditorExtensions\" />\r\n    </xs:appinfo>\r\n  </xs:annotation>\r\n  <xs:element name=\"OrderFile\">\r\n    <xs:annotation>\r\n      <xs:appinfo>\r\n        <b:recordInfo structure=\"delimited\" preserve_delimiter_for_empty_data=\"true\" suppress_trailing_delimiters=\"false\" sequence_number=\"1\" />\r\n      </xs:appinfo>\r\n    </xs:annotation>\r\n    <xs:complexType>\r\n      <xs:sequence>\r\n        <xs:annotation>\r\n          <xs:appinfo>\r\n            <b:groupInfo sequence_number=\"0\" />\r\n          </xs:appinfo>\r\n        </xs:annotation>\r\n        <xs:element name=\"Order\">\r\n          <xs:annotation>\r\n            <xs:appinfo>\r\n              <b:recordInfo sequence_number=\"1\" structure=\"delimited\" preserve_delimiter_for_empty_data=\"true\" suppress_trailing_delimiters=\"false\" child_delimiter_type=\"hex\" child_delimiter=\"0x0D 0x0A\" child_order=\"infix\" />\r\n            </xs:appinfo>\r\n          </xs:annotation>\r\n          <xs:complexType>\r\n            <xs:sequence>\r\n              <xs:annotation>\r\n                <xs:appinfo>\r\n                  <b:groupInfo sequence_number=\"0\" />\r\n                </xs:appinfo>\r\n              </xs:annotation>\r\n              <xs:element name=\"Header\">\r\n                <xs:annotation>\r\n                  <xs:appinfo>\r\n                    <b:recordInfo sequence_number=\"1\" structure=\"delimited\" preserve_delimiter_for_empty_data=\"true\" suppress_trailing_delimiters=\"false\" child_delimiter_type=\"char\" child_delimiter=\"|\" child_order=\"infix\" tag_name=\"HDR|\" />\r\n                  </xs:appinfo>\r\n                </xs:annotation>\r\n                <xs:complexType>\r\n                  <xs:sequence>\r\n                    <xs:annotation>\r\n                      <xs:appinfo>\r\n                        <b:groupInfo sequence_number=\"0\" />\r\n                      </xs:appinfo>\r\n                    </xs:annotation>\r\n                    <xs:element name=\"PODate\" type=\"xs:string\">\r\n                      <xs:annotation>\r\n                        <xs:appinfo>\r\n                          <b:fieldInfo sequence_number=\"1\" justification=\"left\" />\r\n                        </xs:appinfo>\r\n                      </xs:annotation>\r\n                    </xs:element>\r\n                    <xs:element name=\"PONumber\" type=\"xs:string\">\r\n                      <xs:annotation>\r\n                        <xs:appinfo>\r\n                          <b:fieldInfo justification=\"left\" sequence_number=\"2\" />\r\n                        </xs:appinfo>\r\n                      </xs:annotation>\r\n                    </xs:element>\r\n                    <xs:element name=\"CustomerID\" type=\"xs:string\">\r\n                      <xs:annotation>\r\n                        <xs:appinfo>\r\n                          <b:fieldInfo sequence_number=\"3\" justification=\"left\" />\r\n                        </xs:appinfo>\r\n                      </xs:annotation>\r\n                    </xs:element>\r\n                    <xs:element name=\"CustomerContactName\" type=\"xs:string\">\r\n                      <xs:annotation>\r\n                        <xs:appinfo>\r\n                          <b:fieldInfo sequence_number=\"4\" justification=\"left\" />\r\n                        </xs:appinfo>\r\n                      </xs:annotation>\r\n                    </xs:element>\r\n                    <xs:element name=\"CustomerContactPhone\" type=\"xs:string\">\r\n                      <xs:annotation>\r\n                        <xs:appinfo>\r\n                          <b:fieldInfo sequence_number=\"5\" justification=\"left\" />\r\n                        </xs:appinfo>\r\n                      </xs:annotation>\r\n                    </xs:element>\r\n                  </xs:sequence>\r\n                </xs:complexType>\r\n              </xs:element>\r\n              <xs:element minOccurs=\"1\" maxOccurs=\"unbounded\" name=\"LineItems\">\r\n                <xs:annotation>\r\n                  <xs:appinfo>\r\n                    <b:recordInfo sequence_number=\"2\" structure=\"delimited\" preserve_delimiter_for_empty_data=\"true\" suppress_trailing_delimiters=\"false\" child_delimiter_type=\"char\" child_delimiter=\"|\" child_order=\"infix\" tag_name=\"DTL|\" />\r\n                  </xs:appinfo>\r\n                </xs:annotation>\r\n                <xs:complexType>\r\n                  <xs:sequence>\r\n                    <xs:annotation>\r\n                      <xs:appinfo>\r\n                        <b:groupInfo sequence_number=\"0\" />\r\n                      </xs:appinfo>\r\n                    </xs:annotation>\r\n                    <xs:element name=\"PONumber\" type=\"xs:string\">\r\n                      <xs:annotation>\r\n                        <xs:appinfo>\r\n                          <b:fieldInfo sequence_number=\"1\" justification=\"left\" />\r\n                        </xs:appinfo>\r\n                      </xs:annotation>\r\n                    </xs:element>\r\n                    <xs:element name=\"ItemOrdered\" type=\"xs:string\">\r\n                      <xs:annotation>\r\n                        <xs:appinfo>\r\n                          <b:fieldInfo sequence_number=\"2\" justification=\"left\" />\r\n                        </xs:appinfo>\r\n                      </xs:annotation>\r\n                    </xs:element>\r\n                    <xs:element name=\"Quantity\" type=\"xs:string\">\r\n                      <xs:annotation>\r\n                        <xs:appinfo>\r\n                          <b:fieldInfo sequence_number=\"3\" justification=\"left\" />\r\n                        </xs:appinfo>\r\n                      </xs:annotation>\r\n                    </xs:element>\r\n                    <xs:element name=\"UOM\" type=\"xs:string\">\r\n                      <xs:annotation>\r\n                        <xs:appinfo>\r\n                          <b:fieldInfo sequence_number=\"4\" justification=\"left\" />\r\n                        </xs:appinfo>\r\n                      </xs:annotation>\r\n                    </xs:element>\r\n                    <xs:element name=\"Price\" type=\"xs:string\">\r\n                      <xs:annotation>\r\n                        <xs:appinfo>\r\n                          <b:fieldInfo sequence_number=\"5\" justification=\"left\" />\r\n                        </xs:appinfo>\r\n                      </xs:annotation>\r\n                    </xs:element>\r\n                    <xs:element name=\"ExtendedPrice\" type=\"xs:string\">\r\n                      <xs:annotation>\r\n                        <xs:appinfo>\r\n                          <b:fieldInfo sequence_number=\"6\" justification=\"left\" />\r\n                        </xs:appinfo>\r\n                      </xs:annotation>\r\n                    </xs:element>\r\n                    <xs:element name=\"Description\" type=\"xs:string\">\r\n                      <xs:annotation>\r\n                        <xs:appinfo>\r\n                          <b:fieldInfo sequence_number=\"7\" justification=\"left\" />\r\n                        </xs:appinfo>\r\n                      </xs:annotation>\r\n                    </xs:element>\r\n                  </xs:sequence>\r\n                </xs:complexType>\r\n              </xs:element>\r\n            </xs:sequence>\r\n          </xs:complexType>\r\n        </xs:element>\r\n      </xs:sequence>\r\n    </xs:complexType>\r\n  </xs:element>\r\n</xs:schema>"),
				ContentType: to.Ptr("application/xml"),
				Metadata:    map[string]interface{}{},
				SchemaType:  to.Ptr(armlogic.SchemaTypeXML),
			},
		},
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}
