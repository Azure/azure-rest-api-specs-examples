using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.CosmosDB;

// Generated from example definition: specification/cosmos-db/resource-manager/Microsoft.DocumentDB/preview/2024-12-01-preview/examples/NetworkSecurityPerimeterConfigurationGet.json
// this example is just showing the usage of "NetworkSecurityPerimeterConfigurations_Get" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this CosmosDBAccountResource created on azure
// for more information of creating CosmosDBAccountResource, please refer to the document of CosmosDBAccountResource
string subscriptionId = "ffffffff-ffff-ffff-ffff-ffffffffffff";
string resourceGroupName = "res4410";
string accountName = "cosmosTest";
ResourceIdentifier cosmosDBAccountResourceId = CosmosDBAccountResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, accountName);
CosmosDBAccountResource cosmosDBAccount = client.GetCosmosDBAccountResource(cosmosDBAccountResourceId);

// get the collection of this NetworkSecurityPerimeterConfigurationResource
NetworkSecurityPerimeterConfigurationCollection collection = cosmosDBAccount.GetNetworkSecurityPerimeterConfigurations();

// invoke the operation
string networkSecurityPerimeterConfigurationName = "dbedb4e0-40e6-4145-81f3-f1314c150774.resourceAssociation1";
NullableResponse<NetworkSecurityPerimeterConfigurationResource> response = await collection.GetIfExistsAsync(networkSecurityPerimeterConfigurationName);
NetworkSecurityPerimeterConfigurationResource result = response.HasValue ? response.Value : null;

if (result == null)
{
    Console.WriteLine("Succeeded with null as result");
}
else
{
    // the variable result is a resource, you could call other operations on this instance as well
    // but just for demo, we get its data from this resource instance
    NetworkSecurityPerimeterConfigurationData resourceData = result.Data;
    // for demo we just print out the id
    Console.WriteLine($"Succeeded on id: {resourceData.Id}");
}
