using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.AppContainers;

// Generated from example definition: specification/app/resource-manager/Microsoft.App/stable/2022-10-01/examples/Revisions_Get.json
// this example is just showing the usage of "ContainerAppsRevisions_GetRevision" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ContainerAppRevisionResource created on azure
// for more information of creating ContainerAppRevisionResource, please refer to the document of ContainerAppRevisionResource
string subscriptionId = "34adfa4f-cedf-4dc0-ba29-b6d1a69ab345";
string resourceGroupName = "rg";
string containerAppName = "testcontainerApp0";
string revisionName = "testcontainerApp0-pjxhsye";
ResourceIdentifier containerAppRevisionResourceId = ContainerAppRevisionResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, containerAppName, revisionName);
ContainerAppRevisionResource containerAppRevision = client.GetContainerAppRevisionResource(containerAppRevisionResourceId);

// invoke the operation
ContainerAppRevisionResource result = await containerAppRevision.GetAsync();

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
ContainerAppRevisionData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
