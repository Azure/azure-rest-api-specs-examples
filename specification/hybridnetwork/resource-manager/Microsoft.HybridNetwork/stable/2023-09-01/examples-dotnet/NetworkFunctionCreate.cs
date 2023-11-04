using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.HybridNetwork;
using Azure.ResourceManager.HybridNetwork.Models;
using Azure.ResourceManager.Resources;

// Generated from example definition: specification/hybridnetwork/resource-manager/Microsoft.HybridNetwork/stable/2023-09-01/examples/NetworkFunctionCreate.json
// this example is just showing the usage of "NetworkFunctions_CreateOrUpdate" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ResourceGroupResource created on azure
// for more information of creating ResourceGroupResource, please refer to the document of ResourceGroupResource
string subscriptionId = "subid";
string resourceGroupName = "rg";
ResourceIdentifier resourceGroupResourceId = ResourceGroupResource.CreateResourceIdentifier(subscriptionId, resourceGroupName);
ResourceGroupResource resourceGroupResource = client.GetResourceGroupResource(resourceGroupResourceId);

// get the collection of this NetworkFunctionResource
NetworkFunctionCollection collection = resourceGroupResource.GetNetworkFunctions();

// invoke the operation
string networkFunctionName = "testNf";
NetworkFunctionData data = new NetworkFunctionData(new AzureLocation("eastus"))
{
    Properties = new NetworkFunctionValueWithoutSecrets()
    {
        DeploymentValues = "{\"releaseName\":\"testReleaseName\",\"namespace\":\"testNamespace\"}",
        NetworkFunctionDefinitionVersionResourceReference = new OpenDeploymentResourceReference()
        {
            Id = new ResourceIdentifier("/subscriptions/subid/resourcegroups/rg/providers/Microsoft.HybridNetwork/publishers/testVendor/networkFunctionDefinitionGroups/testnetworkFunctionDefinitionGroupName/networkFunctionDefinitionVersions/1.0.1"),
        },
        NfviType = NfviType.AzureArcKubernetes,
        NfviId = new ResourceIdentifier("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/testResourceGroup/providers/Microsoft.ExtendedLocation/customLocations/testCustomLocation"),
        AllowSoftwareUpdate = false,
        RoleOverrideValues =
        {
        "{\"name\":\"testRoleOne\",\"deployParametersMappingRuleProfile\":{\"helmMappingRuleProfile\":{\"helmPackageVersion\":\"2.1.3\",\"values\":\"{\\\"roleOneParam\\\":\\\"roleOneOverrideValue\\\"}\"}}}","{\"name\":\"testRoleTwo\",\"deployParametersMappingRuleProfile\":{\"helmMappingRuleProfile\":{\"releaseName\":\"overrideReleaseName\",\"releaseNamespace\":\"overrideNamespace\",\"values\":\"{\\\"roleTwoParam\\\":\\\"roleTwoOverrideValue\\\"}\"}}}"
        },
    },
};
ArmOperation<NetworkFunctionResource> lro = await collection.CreateOrUpdateAsync(WaitUntil.Completed, networkFunctionName, data);
NetworkFunctionResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
NetworkFunctionData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
