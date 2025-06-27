using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.PostgreSql.FlexibleServers.Models;
using Azure.ResourceManager.PostgreSql.FlexibleServers;

// Generated from example definition: specification/postgresql/resource-manager/Microsoft.DBforPostgreSQL/stable/2024-08-01/examples/VirtualEndpointDelete.json
// this example is just showing the usage of "VirtualEndpoints_Delete" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this VirtualEndpointResource created on azure
// for more information of creating VirtualEndpointResource, please refer to the document of VirtualEndpointResource
string subscriptionId = "ffffffff-ffff-ffff-ffff-ffffffffffff";
string resourceGroupName = "testrg";
string serverName = "pgtestsvc4";
string virtualEndpointName = "pgVirtualEndpoint1";
ResourceIdentifier virtualEndpointResourceId = VirtualEndpointResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, serverName, virtualEndpointName);
VirtualEndpointResource virtualEndpointResource = client.GetVirtualEndpointResource(virtualEndpointResourceId);

// invoke the operation
await virtualEndpointResource.DeleteAsync(WaitUntil.Completed);

Console.WriteLine("Succeeded");
