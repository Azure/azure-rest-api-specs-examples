package armhybridnetwork_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/hybridnetwork/armhybridnetwork/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/db9788dde7a0c2c0d82e4fdf5f7b4de3843937e3/specification/hybridnetwork/resource-manager/Microsoft.HybridNetwork/stable/2023-09-01/examples/NetworkFunctionCreate.json
func ExampleNetworkFunctionsClient_BeginCreateOrUpdate_createNetworkFunctionResource() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armhybridnetwork.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewNetworkFunctionsClient().BeginCreateOrUpdate(ctx, "rg", "testNf", armhybridnetwork.NetworkFunction{
		Location: to.Ptr("eastus"),
		Properties: &armhybridnetwork.NetworkFunctionValueWithoutSecrets{
			AllowSoftwareUpdate: to.Ptr(false),
			ConfigurationType:   to.Ptr(armhybridnetwork.NetworkFunctionConfigurationTypeOpen),
			NetworkFunctionDefinitionVersionResourceReference: &armhybridnetwork.OpenDeploymentResourceReference{
				IDType: to.Ptr(armhybridnetwork.IDTypeOpen),
				ID:     to.Ptr("/subscriptions/subid/resourcegroups/rg/providers/Microsoft.HybridNetwork/publishers/testVendor/networkFunctionDefinitionGroups/testnetworkFunctionDefinitionGroupName/networkFunctionDefinitionVersions/1.0.1"),
			},
			NfviID:   to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/testResourceGroup/providers/Microsoft.ExtendedLocation/customLocations/testCustomLocation"),
			NfviType: to.Ptr(armhybridnetwork.NFVITypeAzureArcKubernetes),
			RoleOverrideValues: []*string{
				to.Ptr("{\"name\":\"testRoleOne\",\"deployParametersMappingRuleProfile\":{\"helmMappingRuleProfile\":{\"helmPackageVersion\":\"2.1.3\",\"values\":\"{\\\"roleOneParam\\\":\\\"roleOneOverrideValue\\\"}\"}}}"),
				to.Ptr("{\"name\":\"testRoleTwo\",\"deployParametersMappingRuleProfile\":{\"helmMappingRuleProfile\":{\"releaseName\":\"overrideReleaseName\",\"releaseNamespace\":\"overrideNamespace\",\"values\":\"{\\\"roleTwoParam\\\":\\\"roleTwoOverrideValue\\\"}\"}}}")},
			DeploymentValues: to.Ptr("{\"releaseName\":\"testReleaseName\",\"namespace\":\"testNamespace\"}"),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.NetworkFunction = armhybridnetwork.NetworkFunction{
	// 	Name: to.Ptr("testNf"),
	// 	Type: to.Ptr("Microsoft.HybridNetwork/networkFunctions"),
	// 	ID: to.Ptr("/subscriptions/subid/resourcegroups/rg/providers/Microsoft.HybridNetwork/networkFunctions/testNf"),
	// 	Location: to.Ptr("eastus"),
	// 	Properties: &armhybridnetwork.NetworkFunctionValueWithoutSecrets{
	// 		AllowSoftwareUpdate: to.Ptr(false),
	// 		ConfigurationType: to.Ptr(armhybridnetwork.NetworkFunctionConfigurationTypeOpen),
	// 		NetworkFunctionDefinitionGroupName: to.Ptr("testnetworkFunctionDefinitionGroupName"),
	// 		NetworkFunctionDefinitionOfferingLocation: to.Ptr("eastus"),
	// 		NetworkFunctionDefinitionVersion: to.Ptr("1.0.1"),
	// 		NetworkFunctionDefinitionVersionResourceReference: &armhybridnetwork.OpenDeploymentResourceReference{
	// 			IDType: to.Ptr(armhybridnetwork.IDTypeOpen),
	// 			ID: to.Ptr("/subscriptions/subid/resourcegroups/rg/providers/Microsoft.HybridNetwork/publishers/testVendor/networkFunctionDefinitionGroups/testnetworkFunctionDefinitionGroupName/networkFunctionDefinitionVersions/1.0.1"),
	// 		},
	// 		NfviID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/testResourceGroup/providers/Microsoft.ExtendedLocation/customLocations/testCustomLocation"),
	// 		NfviType: to.Ptr(armhybridnetwork.NFVITypeAzureArcKubernetes),
	// 		ProvisioningState: to.Ptr(armhybridnetwork.ProvisioningStateSucceeded),
	// 		PublisherName: to.Ptr("testVendor"),
	// 		PublisherScope: to.Ptr(armhybridnetwork.PublisherScope("Public")),
	// 		RoleOverrideValues: []*string{
	// 			to.Ptr("{\"name\":\"testRoleOne\",\"deployParametersMappingRuleProfile\":{\"helmMappingRuleProfile\":{\"helmPackageVersion\":\"2.1.3\",\"values\":\"{\\\"roleOneParam\\\":\\\"roleOneOverrideValue\\\"}\"}}}"),
	// 			to.Ptr("{\"name\":\"testRoleTwo\",\"deployParametersMappingRuleProfile\":{\"helmMappingRuleProfile\":{\"releaseName\":\"overrideReleaseName\",\"releaseNamespace\":\"overrideNamespace\",\"values\":\"{\\\"roleTwoParam\\\":\\\"roleTwoOverrideValue\\\"}\"}}}")},
	// 			DeploymentValues: to.Ptr("{\"releaseName\":\"testReleaseName\",\"namespace\":\"testNamespace\"}"),
	// 		},
	// 	}
}
