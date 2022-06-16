package armcontainerservice_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/containerservice/armcontainerservice"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/containerservice/resource-manager/Microsoft.ContainerService/stable/2022-03-01/examples/MaintenanceConfigurationsCreate_Update.json
func ExampleMaintenanceConfigurationsClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
		return
	}
	ctx := context.Background()
	client, err := armcontainerservice.NewMaintenanceConfigurationsClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
		return
	}
	res, err := client.CreateOrUpdate(ctx,
		"<resource-group-name>",
		"<resource-name>",
		"<config-name>",
		armcontainerservice.MaintenanceConfiguration{
			Properties: &armcontainerservice.MaintenanceConfigurationProperties{
				NotAllowedTime: []*armcontainerservice.TimeSpan{
					{
						End:   to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-11-30T12:00:00Z"); return t }()),
						Start: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-11-26T03:00:00Z"); return t }()),
					}},
				TimeInWeek: []*armcontainerservice.TimeInWeek{
					{
						Day: to.Ptr(armcontainerservice.WeekDayMonday),
						HourSlots: []*int32{
							to.Ptr[int32](1),
							to.Ptr[int32](2)},
					}},
			},
		},
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
		return
	}
	// TODO: use response item
	_ = res
}
