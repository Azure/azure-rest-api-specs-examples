using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.AppContainers;

// Generated from example definition: specification/app/resource-manager/Microsoft.App/stable/2023-05-01/examples/Replicas_Get.json
// this example is just showing the usage of "ContainerAppsRevisionReplicas_GetReplica" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ContainerAppRevisionResource created on azure
// for more information of creating ContainerAppRevisionResource, please refer to the document of ContainerAppRevisionResource
string subscriptionId = "651f8027-33e8-4ec4-97b4-f6e9f3dc8744";
string resourceGroupName = "workerapps-rg-xj";
string containerAppName = "myapp";
string revisionName = "myapp--0wlqy09";
ResourceIdentifier containerAppRevisionResourceId = ContainerAppRevisionResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, containerAppName, revisionName);
ContainerAppRevisionResource containerAppRevision = client.GetContainerAppRevisionResource(containerAppRevisionResourceId);

// get the collection of this ContainerAppReplicaResource
ContainerAppReplicaCollection collection = containerAppRevision.GetContainerAppReplicas();

// invoke the operation
string replicaName = "myapp--0wlqy09-5d9774cff-5wnd8";
NullableResponse<ContainerAppReplicaResource> response = await collection.GetIfExistsAsync(replicaName);
ContainerAppReplicaResource result = response.HasValue ? response.Value : null;

if (result == null)
{
    Console.WriteLine($"Succeeded with null as result");
}
else
{
    // the variable result is a resource, you could call other operations on this instance as well
    // but just for demo, we get its data from this resource instance
    ContainerAppReplicaData resourceData = result.Data;
    // for demo we just print out the id
    Console.WriteLine($"Succeeded on id: {resourceData.Id}");
}
