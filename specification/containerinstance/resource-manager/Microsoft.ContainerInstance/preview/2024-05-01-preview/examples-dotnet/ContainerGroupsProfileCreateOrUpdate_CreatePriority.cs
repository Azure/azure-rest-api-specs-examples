using Azure;
using Azure.ResourceManager;
using System;
using System.Collections.Generic;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.ContainerInstance.Models;
using Azure.ResourceManager.Resources;
using Azure.ResourceManager.ContainerInstance;

// Generated from example definition: specification/containerinstance/resource-manager/Microsoft.ContainerInstance/preview/2024-05-01-preview/examples/ContainerGroupsProfileCreateOrUpdate_CreatePriority.json
// this example is just showing the usage of "ContainerGroupProfiles_CreateOrUpdate" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ResourceGroupResource created on azure
// for more information of creating ResourceGroupResource, please refer to the document of ResourceGroupResource
string subscriptionId = "00000000-0000-0000-0000-000000000000";
string resourceGroupName = "demo";
ResourceIdentifier resourceGroupResourceId = ResourceGroupResource.CreateResourceIdentifier(subscriptionId, resourceGroupName);
ResourceGroupResource resourceGroupResource = client.GetResourceGroupResource(resourceGroupResourceId);

// get the collection of this ContainerGroupProfileResource
ContainerGroupProfileCollection collection = resourceGroupResource.GetContainerGroupProfiles();

// invoke the operation
string containerGroupProfileName = "demo1";
ContainerGroupProfileData data = new ContainerGroupProfileData(new AzureLocation("eastus"), new ContainerInstanceContainer[]
{
new ContainerInstanceContainer("test-container-001")
{
Image = "alpine:latest",
Command =
{
"/bin/sh","-c","sleep 10"
},
Resources = new ContainerResourceRequirements(new ContainerResourceRequestsContent(1,1)),
}
}, ContainerInstanceOperatingSystemType.Linux)
{
    RestartPolicy = ContainerGroupRestartPolicy.Never,
    Sku = ContainerGroupSku.Standard,
    Priority = ContainerGroupPriority.Spot,
};
ArmOperation<ContainerGroupProfileResource> lro = await collection.CreateOrUpdateAsync(WaitUntil.Completed, containerGroupProfileName, data);
ContainerGroupProfileResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
ContainerGroupProfileData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
