using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Grafana;

// Generated from example definition: 2024-11-01-preview/ManagedPrivateEndpoints_Get.json
// this example is just showing the usage of "ManagedPrivateEndpointModel_Get" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ManagedGrafanaResource created on azure
// for more information of creating ManagedGrafanaResource, please refer to the document of ManagedGrafanaResource
string subscriptionId = "00000000-0000-0000-0000-000000000000";
string resourceGroupName = "myResourceGroup";
string workspaceName = "myWorkspace";
ResourceIdentifier managedGrafanaResourceId = ManagedGrafanaResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, workspaceName);
ManagedGrafanaResource managedGrafana = client.GetManagedGrafanaResource(managedGrafanaResourceId);

// get the collection of this ManagedPrivateEndpointModelResource
ManagedPrivateEndpointModelCollection collection = managedGrafana.GetManagedPrivateEndpointModels();

// invoke the operation
string managedPrivateEndpointName = "myMPEName";
NullableResponse<ManagedPrivateEndpointModelResource> response = await collection.GetIfExistsAsync(managedPrivateEndpointName);
ManagedPrivateEndpointModelResource result = response.HasValue ? response.Value : null;

if (result == null)
{
    Console.WriteLine("Succeeded with null as result");
}
else
{
    // the variable result is a resource, you could call other operations on this instance as well
    // but just for demo, we get its data from this resource instance
    ManagedPrivateEndpointModelData resourceData = result.Data;
    // for demo we just print out the id
    Console.WriteLine($"Succeeded on id: {resourceData.Id}");
}
