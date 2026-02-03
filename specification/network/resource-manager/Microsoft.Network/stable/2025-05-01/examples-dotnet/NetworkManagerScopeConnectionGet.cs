using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Network;

// Generated from example definition: specification/network/resource-manager/Microsoft.Network/stable/2025-05-01/examples/NetworkManagerScopeConnectionGet.json
// this example is just showing the usage of "ScopeConnections_Get" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ScopeConnectionResource created on azure
// for more information of creating ScopeConnectionResource, please refer to the document of ScopeConnectionResource
string subscriptionId = "00000000-0000-0000-0000-000000000000";
string resourceGroupName = "rg1";
string networkManagerName = "testNetworkManager";
string scopeConnectionName = "TestScopeConnection";
ResourceIdentifier scopeConnectionResourceId = ScopeConnectionResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, networkManagerName, scopeConnectionName);
ScopeConnectionResource scopeConnection = client.GetScopeConnectionResource(scopeConnectionResourceId);

// invoke the operation
ScopeConnectionResource result = await scopeConnection.GetAsync();

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
ScopeConnectionData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
