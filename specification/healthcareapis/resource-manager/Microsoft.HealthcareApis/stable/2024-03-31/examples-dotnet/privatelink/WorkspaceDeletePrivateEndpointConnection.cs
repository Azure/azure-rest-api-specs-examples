using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.HealthcareApis.Models;
using Azure.ResourceManager.HealthcareApis;

// Generated from example definition: specification/healthcareapis/resource-manager/Microsoft.HealthcareApis/stable/2024-03-31/examples/privatelink/WorkspaceDeletePrivateEndpointConnection.json
// this example is just showing the usage of "WorkspacePrivateEndpointConnections_Delete" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this HealthcareApisWorkspacePrivateEndpointConnectionResource created on azure
// for more information of creating HealthcareApisWorkspacePrivateEndpointConnectionResource, please refer to the document of HealthcareApisWorkspacePrivateEndpointConnectionResource
string subscriptionId = "subid";
string resourceGroupName = "testRG";
string workspaceName = "workspace1";
string privateEndpointConnectionName = "myConnection";
ResourceIdentifier healthcareApisWorkspacePrivateEndpointConnectionResourceId = HealthcareApisWorkspacePrivateEndpointConnectionResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, workspaceName, privateEndpointConnectionName);
HealthcareApisWorkspacePrivateEndpointConnectionResource healthcareApisWorkspacePrivateEndpointConnection = client.GetHealthcareApisWorkspacePrivateEndpointConnectionResource(healthcareApisWorkspacePrivateEndpointConnectionResourceId);

// invoke the operation
await healthcareApisWorkspacePrivateEndpointConnection.DeleteAsync(WaitUntil.Completed);

Console.WriteLine($"Succeeded");
