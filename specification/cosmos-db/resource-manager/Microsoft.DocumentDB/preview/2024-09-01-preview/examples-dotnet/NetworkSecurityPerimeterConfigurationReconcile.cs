using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.CosmosDB;

// Generated from example definition: specification/cosmos-db/resource-manager/Microsoft.DocumentDB/preview/2024-09-01-preview/examples/NetworkSecurityPerimeterConfigurationReconcile.json
// this example is just showing the usage of "NetworkSecurityPerimeterConfigurations_Reconcile" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this NetworkSecurityPerimeterConfigurationResource created on azure
// for more information of creating NetworkSecurityPerimeterConfigurationResource, please refer to the document of NetworkSecurityPerimeterConfigurationResource
string subscriptionId = "ffffffff-ffff-ffff-ffff-ffffffffffff";
string resourceGroupName = "res4410";
string accountName = "sto8607";
string networkSecurityPerimeterConfigurationName = "dbedb4e0-40e6-4145-81f3-f1314c150774.resourceAssociation1";
ResourceIdentifier networkSecurityPerimeterConfigurationResourceId = NetworkSecurityPerimeterConfigurationResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, accountName, networkSecurityPerimeterConfigurationName);
NetworkSecurityPerimeterConfigurationResource networkSecurityPerimeterConfiguration = client.GetNetworkSecurityPerimeterConfigurationResource(networkSecurityPerimeterConfigurationResourceId);

// invoke the operation
await networkSecurityPerimeterConfiguration.ReconcileAsync(WaitUntil.Completed);

Console.WriteLine($"Succeeded");
