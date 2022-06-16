package armaad_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/aad/armaad"
)

// x-ms-original-file: specification/azureactivedirectory/resource-manager/Microsoft.Aadiam/preview/2017-04-01-preview/examples/createOrUpdateDiagnosticSetting.json
func ExampleDiagnosticSettingsClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armaad.NewDiagnosticSettingsClient(cred, nil)
	res, err := client.CreateOrUpdate(ctx,
		"<name>",
		armaad.DiagnosticSettingsResource{
			Properties: &armaad.DiagnosticSettings{
				EventHubAuthorizationRuleID: to.StringPtr("<event-hub-authorization-rule-id>"),
				EventHubName:                to.StringPtr("<event-hub-name>"),
				Logs: []*armaad.LogSettings{
					{
						Category: armaad.Category("AuditLogs").ToPtr(),
						Enabled:  to.BoolPtr(true),
						RetentionPolicy: &armaad.RetentionPolicy{
							Days:    to.Int32Ptr(0),
							Enabled: to.BoolPtr(false),
						},
					}},
				StorageAccountID: to.StringPtr("<storage-account-id>"),
				WorkspaceID:      to.StringPtr("<workspace-id>"),
			},
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.DiagnosticSettingsClientCreateOrUpdateResult)
}
