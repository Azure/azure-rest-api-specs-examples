using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.MySql.FlexibleServers;

// Generated from example definition: 2024-12-30/PrivateEndpointConnectionDelete.json
// this example is just showing the usage of "PrivateEndpointConnection_Delete" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this MySqlFlexibleServersPrivateEndpointConnectionResource created on azure
// for more information of creating MySqlFlexibleServersPrivateEndpointConnectionResource, please refer to the document of MySqlFlexibleServersPrivateEndpointConnectionResource
string subscriptionId = "00000000-1111-2222-3333-444444444444";
string resourceGroupName = "Default";
string serverName = "test-svr";
string privateEndpointConnectionName = "private-endpoint-connection-name";
ResourceIdentifier mySqlFlexibleServersPrivateEndpointConnectionResourceId = MySqlFlexibleServersPrivateEndpointConnectionResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, serverName, privateEndpointConnectionName);
MySqlFlexibleServersPrivateEndpointConnectionResource mySqlFlexibleServersPrivateEndpointConnection = client.GetMySqlFlexibleServersPrivateEndpointConnectionResource(mySqlFlexibleServersPrivateEndpointConnectionResourceId);

// invoke the operation
await mySqlFlexibleServersPrivateEndpointConnection.DeleteAsync(WaitUntil.Completed);

Console.WriteLine("Succeeded");
