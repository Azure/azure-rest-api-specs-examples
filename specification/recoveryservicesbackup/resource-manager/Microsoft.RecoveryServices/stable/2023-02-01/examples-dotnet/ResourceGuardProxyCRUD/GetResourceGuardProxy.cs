using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.RecoveryServicesBackup;
using Azure.ResourceManager.RecoveryServicesBackup.Models;

// Generated from example definition: specification/recoveryservicesbackup/resource-manager/Microsoft.RecoveryServices/stable/2023-02-01/examples/ResourceGuardProxyCRUD/GetResourceGuardProxy.json
// this example is just showing the usage of "ResourceGuardProxy_Get" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ResourceGuardProxyResource created on azure
// for more information of creating ResourceGuardProxyResource, please refer to the document of ResourceGuardProxyResource
string subscriptionId = "0b352192-dcac-4cc7-992e-a96190ccc68c";
string resourceGroupName = "SampleResourceGroup";
string vaultName = "sampleVault";
string resourceGuardProxyName = "swaggerExample";
ResourceIdentifier resourceGuardProxyResourceId = ResourceGuardProxyResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, vaultName, resourceGuardProxyName);
ResourceGuardProxyResource resourceGuardProxy = client.GetResourceGuardProxyResource(resourceGuardProxyResourceId);

// invoke the operation
ResourceGuardProxyResource result = await resourceGuardProxy.GetAsync();

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
ResourceGuardProxyData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
