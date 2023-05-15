using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.AppContainers;
using Azure.ResourceManager.AppContainers.Models;

// Generated from example definition: specification/app/resource-manager/Microsoft.App/preview/2022-11-01-preview/examples/ConnectedEnvironmentsCertificates_ListByConnectedEnvironment.json
// this example is just showing the usage of "ConnectedEnvironmentsCertificates_List" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ContainerAppConnectedEnvironmentResource created on azure
// for more information of creating ContainerAppConnectedEnvironmentResource, please refer to the document of ContainerAppConnectedEnvironmentResource
string subscriptionId = "34adfa4f-cedf-4dc0-ba29-b6d1a69ab345";
string resourceGroupName = "examplerg";
string connectedEnvironmentName = "testcontainerenv";
ResourceIdentifier containerAppConnectedEnvironmentResourceId = ContainerAppConnectedEnvironmentResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, connectedEnvironmentName);
ContainerAppConnectedEnvironmentResource containerAppConnectedEnvironment = client.GetContainerAppConnectedEnvironmentResource(containerAppConnectedEnvironmentResourceId);

// get the collection of this ContainerAppConnectedEnvironmentCertificateResource
ContainerAppConnectedEnvironmentCertificateCollection collection = containerAppConnectedEnvironment.GetContainerAppConnectedEnvironmentCertificates();

// invoke the operation and iterate over the result
await foreach (ContainerAppConnectedEnvironmentCertificateResource item in collection.GetAllAsync())
{
    // the variable item is a resource, you could call other operations on this instance as well
    // but just for demo, we get its data from this resource instance
    ContainerAppCertificateData resourceData = item.Data;
    // for demo we just print out the id
    Console.WriteLine($"Succeeded on id: {resourceData.Id}");
}

Console.WriteLine($"Succeeded");
