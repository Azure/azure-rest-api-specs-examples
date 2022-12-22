package armredhatopenshift_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/redhatopenshift/armredhatopenshift"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/redhatopenshift/resource-manager/Microsoft.RedHatOpenShift/stable/2022-09-04/examples/MachinePools_CreateOrUpdate.json
func ExampleMachinePoolsClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armredhatopenshift.NewMachinePoolsClient("subscriptionId", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.CreateOrUpdate(ctx, "resourceGroup", "resourceName", "childResourceName", armredhatopenshift.MachinePool{
		Properties: &armredhatopenshift.MachinePoolProperties{
			Resources: to.Ptr("ewogICAgImFwaVZlcnNpb24iOiAiaGl2ZS5vcGVuc2hpZnQuaW8vdjEiLAogICAgImtpbmQiOiAiTWFjaGluZVBvb2wiLAogICAgIm1ldGFkYXRhIjogewogICAgICAgICJuYW1lIjogInRlc3QtY2x1c3Rlci13b3JrZXIiLAogICAgICAgICJuYW1lc3BhY2UiOiAiYXJvLWY2MGFlOGEyLWJjYTEtNDk4Ny05MDU2LVhYWFhYWFhYWFhYWCIKICAgIH0sCiAgICAic3BlYyI6IHsKICAgICAgICAiY2x1c3RlckRlcGxveW1lbnRSZWYiOiB7CiAgICAgICAgICAgICJuYW1lIjogInRlc3QtY2x1c3RlciIKICAgICAgICB9LAogICAgICAgICJuYW1lIjogIndvcmtlciIsCiAgICAgICAgInBsYXRmb3JtIjogewogICAgICAgICAgICAiYXdzIjogewogICAgICAgICAgICAgICAgInJvb3RWb2x1bWUiOiB7CiAgICAgICAgICAgICAgICAgICAgImlvcHMiOiAwLAogICAgICAgICAgICAgICAgICAgICJzaXplIjogMzAwLAogICAgICAgICAgICAgICAgICAgICJ0eXBlIjogImdwMyIKICAgICAgICAgICAgICAgIH0sCiAgICAgICAgICAgICAgICAidHlwZSI6ICJtNS54bGFyZ2UiLAogICAgICAgICAgICAgICAgInpvbmVzIjogWwogICAgICAgICAgICAgICAgICAgICJ1cy1lYXN0LTFhIgogICAgICAgICAgICAgICAgXQogICAgICAgICAgICB9CiAgICAgICAgfSwKICAgICAgICAicmVwbGljYXMiOiAyCiAgICB9LAogICAgInN0YXR1cyI6IHsKICAgICAgICAiY29uZGl0aW9ucyI6IFsKICAgICAgICBdCiAgICB9Cn0K"),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}
