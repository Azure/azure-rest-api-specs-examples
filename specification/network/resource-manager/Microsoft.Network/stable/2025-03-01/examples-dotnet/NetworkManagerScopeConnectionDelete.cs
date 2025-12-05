using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Network;

// Generated from example definition: specification/network/resource-manager/Microsoft.Network/stable/2025-03-01/examples/NetworkManagerScopeConnectionDelete.json
// this example is just showing the usage of "ScopeConnections_Delete" operation, for the dependent resources, they will have to be created separately.

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
await scopeConnection.DeleteAsync(WaitUntil.Completed);

Console.WriteLine("Succeeded");
