using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Grafana;

// Generated from example definition: specification/dashboard/resource-manager/Microsoft.Dashboard/stable/2023-09-01/examples/PrivateEndpointConnections_Delete.json
// this example is just showing the usage of "PrivateEndpointConnections_Delete" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this GrafanaPrivateEndpointConnectionResource created on azure
// for more information of creating GrafanaPrivateEndpointConnectionResource, please refer to the document of GrafanaPrivateEndpointConnectionResource
string subscriptionId = "00000000-0000-0000-0000-000000000000";
string resourceGroupName = "myResourceGroup";
string workspaceName = "myWorkspace";
string privateEndpointConnectionName = "myConnection";
ResourceIdentifier grafanaPrivateEndpointConnectionResourceId = GrafanaPrivateEndpointConnectionResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, workspaceName, privateEndpointConnectionName);
GrafanaPrivateEndpointConnectionResource grafanaPrivateEndpointConnection = client.GetGrafanaPrivateEndpointConnectionResource(grafanaPrivateEndpointConnectionResourceId);

// invoke the operation
await grafanaPrivateEndpointConnection.DeleteAsync(WaitUntil.Completed);

Console.WriteLine($"Succeeded");
