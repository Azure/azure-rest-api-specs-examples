using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.ManagementGroups;
using Azure.ResourceManager.Network;

// Generated from example definition: specification/network/resource-manager/Microsoft.Network/stable/2023-05-01/examples/NetworkManagerConnectionManagementGroupGet.json
// this example is just showing the usage of "ManagementGroupNetworkManagerConnections_Get" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ManagementGroupResource created on azure
// for more information of creating ManagementGroupResource, please refer to the document of ManagementGroupResource
string managementGroupId = "managementGroupA";
ResourceIdentifier managementGroupResourceId = ManagementGroupResource.CreateResourceIdentifier(managementGroupId);
ManagementGroupResource managementGroupResource = client.GetManagementGroupResource(managementGroupResourceId);

// get the collection of this ManagementGroupNetworkManagerConnectionResource
ManagementGroupNetworkManagerConnectionCollection collection = managementGroupResource.GetManagementGroupNetworkManagerConnections();

// invoke the operation
string networkManagerConnectionName = "TestNMConnection";
bool result = await collection.ExistsAsync(networkManagerConnectionName);

Console.WriteLine($"Succeeded: {result}");
