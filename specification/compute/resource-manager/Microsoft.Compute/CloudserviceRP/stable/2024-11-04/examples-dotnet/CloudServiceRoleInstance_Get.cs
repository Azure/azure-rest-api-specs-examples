using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Compute;

// Generated from example definition: specification/compute/resource-manager/Microsoft.Compute/CloudserviceRP/stable/2024-11-04/examples/CloudServiceRoleInstance_Get.json
// this example is just showing the usage of "CloudServiceRoleInstances_Get" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this CloudServiceResource created on azure
// for more information of creating CloudServiceResource, please refer to the document of CloudServiceResource
string subscriptionId = "{subscription-id}";
string resourceGroupName = "ConstosoRG";
string cloudServiceName = "{cs-name}";
ResourceIdentifier cloudServiceResourceId = CloudServiceResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, cloudServiceName);
CloudServiceResource cloudService = client.GetCloudServiceResource(cloudServiceResourceId);

// get the collection of this CloudServiceRoleInstanceResource
CloudServiceRoleInstanceCollection collection = cloudService.GetCloudServiceRoleInstances();

// invoke the operation
string roleInstanceName = "{roleInstance-name}";
NullableResponse<CloudServiceRoleInstanceResource> response = await collection.GetIfExistsAsync(roleInstanceName);
CloudServiceRoleInstanceResource result = response.HasValue ? response.Value : null;

if (result == null)
{
    Console.WriteLine("Succeeded with null as result");
}
else
{
    // the variable result is a resource, you could call other operations on this instance as well
    // but just for demo, we get its data from this resource instance
    CloudServiceRoleInstanceData resourceData = result.Data;
    // for demo we just print out the id
    Console.WriteLine($"Succeeded on id: {resourceData.Id}");
}
