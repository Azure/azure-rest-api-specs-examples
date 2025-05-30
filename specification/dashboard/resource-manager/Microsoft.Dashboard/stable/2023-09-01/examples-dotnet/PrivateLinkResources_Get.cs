using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Grafana;

// Generated from example definition: specification/dashboard/resource-manager/Microsoft.Dashboard/stable/2023-09-01/examples/PrivateLinkResources_Get.json
// this example is just showing the usage of "PrivateLinkResources_Get" operation, for the dependent resources, they will have to be created separately.

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

// get the collection of this GrafanaPrivateLinkResource
GrafanaPrivateLinkResourceCollection collection = managedGrafana.GetGrafanaPrivateLinkResources();

// invoke the operation
string privateLinkResourceName = "grafana";
NullableResponse<GrafanaPrivateLinkResource> response = await collection.GetIfExistsAsync(privateLinkResourceName);
GrafanaPrivateLinkResource result = response.HasValue ? response.Value : null;

if (result == null)
{
    Console.WriteLine("Succeeded with null as result");
}
else
{
    // the variable result is a resource, you could call other operations on this instance as well
    // but just for demo, we get its data from this resource instance
    GrafanaPrivateLinkResourceData resourceData = result.Data;
    // for demo we just print out the id
    Console.WriteLine($"Succeeded on id: {resourceData.Id}");
}
