package armdeploymentmanager_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/deploymentmanager/armdeploymentmanager"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/7a2ac91de424f271cf91cc8009f3fe9ee8249086/specification/deploymentmanager/resource-manager/Microsoft.DeploymentManager/preview/2019-11-01-preview/examples/serviceunit_createorupdate.json
func ExampleServiceUnitsClient_BeginCreateOrUpdate_createServiceUnitUsingRelativePathsIntoTheArtifactSource() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdeploymentmanager.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewServiceUnitsClient().BeginCreateOrUpdate(ctx, "myResourceGroup", "myTopology", "myService", "myServiceUnit", armdeploymentmanager.ServiceUnitResource{
		Location: to.Ptr("centralus"),
		Tags:     map[string]*string{},
		Properties: &armdeploymentmanager.ServiceUnitResourceProperties{
			Artifacts: &armdeploymentmanager.ServiceUnitArtifacts{
				ParametersArtifactSourceRelativePath: to.Ptr("parameter/myTopologyUnit.parameters.json"),
				TemplateArtifactSourceRelativePath:   to.Ptr("templates/myTopologyUnit.template.json"),
			},
			DeploymentMode:      to.Ptr(armdeploymentmanager.DeploymentModeIncremental),
			TargetResourceGroup: to.Ptr("myDeploymentResourceGroup"),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	_, err = poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
}
