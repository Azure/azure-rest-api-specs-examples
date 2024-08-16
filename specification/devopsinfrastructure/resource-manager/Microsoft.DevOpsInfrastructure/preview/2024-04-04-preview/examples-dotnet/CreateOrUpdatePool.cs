using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.DevOpsInfrastructure.Models;
using Azure.ResourceManager.Resources;
using Azure.ResourceManager.DevOpsInfrastructure;

// Generated from example definition: specification/devopsinfrastructure/resource-manager/Microsoft.DevOpsInfrastructure/preview/2024-04-04-preview/examples/CreateOrUpdatePool.json
// this example is just showing the usage of "Pools_CreateOrUpdate" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ResourceGroupResource created on azure
// for more information of creating ResourceGroupResource, please refer to the document of ResourceGroupResource
string subscriptionId = "a2e95d27-c161-4b61-bda4-11512c14c2c2";
string resourceGroupName = "rg";
ResourceIdentifier resourceGroupResourceId = ResourceGroupResource.CreateResourceIdentifier(subscriptionId, resourceGroupName);
ResourceGroupResource resourceGroupResource = client.GetResourceGroupResource(resourceGroupResourceId);

// get the collection of this DevOpsPoolResource
DevOpsPoolCollection collection = resourceGroupResource.GetDevOpsPools();

// invoke the operation
string poolName = "pool";
DevOpsPoolData data = new DevOpsPoolData(new AzureLocation("eastus"))
{
    Properties = new DevOpsPoolProperties(10, new DevOpsAzureOrganizationProfile(new DevOpsOrganization[]
{
new DevOpsOrganization(new Uri("https://mseng.visualstudio.com"))
}), new DevOpsStatelessAgentProfile(), new DevOpsVmssFabricProfile(new DevOpsAzureSku("Standard_D4ads_v5"), new DevOpsPoolVmImage[]
{
new DevOpsPoolVmImage()
{
ResourceId = "/MicrosoftWindowsServer/WindowsServer/2019-Datacenter/latest",
}
}), "/subscriptions/222e81d0-cf38-4dab-baa5-289bf16baaa4/resourceGroups/rg-1es-devcenter/providers/Microsoft.DevCenter/projects/1ES")
    {
        ProvisioningState = DevOpsInfrastructureProvisioningState.Succeeded,
    },
};
ArmOperation<DevOpsPoolResource> lro = await collection.CreateOrUpdateAsync(WaitUntil.Completed, poolName, data);
DevOpsPoolResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
DevOpsPoolData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
