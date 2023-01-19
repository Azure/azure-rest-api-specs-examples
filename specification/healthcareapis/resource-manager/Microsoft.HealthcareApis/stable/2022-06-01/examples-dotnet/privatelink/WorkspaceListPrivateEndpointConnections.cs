using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.HealthcareApis;
using Azure.ResourceManager.HealthcareApis.Models;

// Generated from example definition: specification/healthcareapis/resource-manager/Microsoft.HealthcareApis/stable/2022-06-01/examples/privatelink/WorkspaceListPrivateEndpointConnections.json
// this example is just showing the usage of "WorkspacePrivateEndpointConnections_ListByWorkspace" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this HealthcareApisWorkspaceResource created on azure
// for more information of creating HealthcareApisWorkspaceResource, please refer to the document of HealthcareApisWorkspaceResource
string subscriptionId = "subid";
string resourceGroupName = "testRG";
string workspaceName = "workspace1";
ResourceIdentifier healthcareApisWorkspaceResourceId = HealthcareApisWorkspaceResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, workspaceName);
HealthcareApisWorkspaceResource healthcareApisWorkspace = client.GetHealthcareApisWorkspaceResource(healthcareApisWorkspaceResourceId);

// get the collection of this HealthcareApisWorkspacePrivateEndpointConnectionResource
HealthcareApisWorkspacePrivateEndpointConnectionCollection collection = healthcareApisWorkspace.GetHealthcareApisWorkspacePrivateEndpointConnections();

// invoke the operation and iterate over the result
await foreach (HealthcareApisWorkspacePrivateEndpointConnectionResource item in collection.GetAllAsync())
{
    // the variable item is a resource, you could call other operations on this instance as well
    // but just for demo, we get its data from this resource instance
    HealthcareApisPrivateEndpointConnectionData resourceData = item.Data;
    // for demo we just print out the id
    Console.WriteLine($"Succeeded on id: {resourceData.Id}");
}

Console.WriteLine($"Succeeded");
