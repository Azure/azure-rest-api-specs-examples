package armsecurityinsight_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/securityinsight/armsecurityinsight"
)

// x-ms-original-file: specification/securityinsights/resource-manager/Microsoft.SecurityInsights/preview/2021-09-01-preview/examples/settings/UpdateEyesOnSetting.json
func ExampleProductSettingsClient_Update() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armsecurityinsight.NewProductSettingsClient("<subscription-id>", cred, nil)
	res, err := client.Update(ctx,
		"<resource-group-name>",
		"<workspace-name>",
		"<settings-name>",
		&armsecurityinsight.EyesOn{
			Etag:       to.StringPtr("<etag>"),
			Kind:       armsecurityinsight.SettingKind("EyesOn").ToPtr(),
			Properties: &armsecurityinsight.EyesOnSettingsProperties{},
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.ProductSettingsClientUpdateResult)
}
