using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.RecoveryServicesBackup;
using Azure.ResourceManager.RecoveryServicesBackup.Models;
using Azure.ResourceManager.Resources;

// Generated from example definition: specification/recoveryservicesbackup/resource-manager/Microsoft.RecoveryServices/stable/2023-04-01/examples/ResourceGuardProxyCRUD/GetResourceGuardProxy.json
// this example is just showing the usage of "ResourceGuardProxy_Get" operation, for the dependent resources, they will have to be created separately.

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
bool result = await collection.ExistsAsync(resourceGuardProxyName);

Console.WriteLine($"Succeeded: {result}");
