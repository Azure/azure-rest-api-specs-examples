using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.RecoveryServicesBackup.Models;
using Azure.ResourceManager.Resources;
using Azure.ResourceManager.RecoveryServicesBackup;

// Generated from example definition: specification/recoveryservicesbackup/resource-manager/Microsoft.RecoveryServices/stable/2023-06-01/examples/ResourceGuardProxyCRUD/PutResourceGuardProxy.json
// this example is just showing the usage of "ResourceGuardProxy_Put" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ResourceGroupResource created on azure
// for more information of creating ResourceGroupResource, please refer to the document of ResourceGroupResource
string subscriptionId = "0b352192-dcac-4cc7-992e-a96190ccc68c";
string resourceGroupName = "SampleResourceGroup";
ResourceIdentifier resourceGroupResourceId = ResourceGroupResource.CreateResourceIdentifier(subscriptionId, resourceGroupName);
ResourceGroupResource resourceGroupResource = client.GetResourceGroupResource(resourceGroupResourceId);

// get the collection of this ResourceGuardProxyResource
string vaultName = "sampleVault";
ResourceGuardProxyCollection collection = resourceGroupResource.GetResourceGuardProxies(vaultName);

// invoke the operation
string resourceGuardProxyName = "swaggerExample";
ResourceGuardProxyData data = new ResourceGuardProxyData(default)
{
    Properties = new ResourceGuardProxyProperties(new ResourceIdentifier("/subscriptions/c999d45b-944f-418c-a0d8-c3fcfd1802c8/resourceGroups/vaultguardRGNew/providers/Microsoft.DataProtection/resourceGuards/VaultGuardTestNew")),
};
ArmOperation<ResourceGuardProxyResource> lro = await collection.CreateOrUpdateAsync(WaitUntil.Completed, resourceGuardProxyName, data);
ResourceGuardProxyResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
ResourceGuardProxyData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
