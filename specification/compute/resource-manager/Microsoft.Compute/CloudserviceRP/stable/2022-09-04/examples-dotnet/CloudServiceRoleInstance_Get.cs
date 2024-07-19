using Azure;
using Azure.ResourceManager;
using System;
using System.IO;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Compute.Models;
using Azure.ResourceManager.Compute;

// Generated from example definition: specification/compute/resource-manager/Microsoft.Compute/CloudserviceRP/stable/2022-09-04/examples/CloudServiceRoleInstance_Get.json
// this example is just showing the usage of "CloudServiceRoleInstances_Get" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this CloudServiceRoleInstanceResource created on azure
// for more information of creating CloudServiceRoleInstanceResource, please refer to the document of CloudServiceRoleInstanceResource
string subscriptionId = "{subscription-id}";
string resourceGroupName = "ConstosoRG";
string cloudServiceName = "{cs-name}";
string roleInstanceName = "{roleInstance-name}";
ResourceIdentifier cloudServiceRoleInstanceResourceId = CloudServiceRoleInstanceResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, cloudServiceName, roleInstanceName);
CloudServiceRoleInstanceResource cloudServiceRoleInstance = client.GetCloudServiceRoleInstanceResource(cloudServiceRoleInstanceResourceId);

// invoke the operation
CloudServiceRoleInstanceResource result = await cloudServiceRoleInstance.GetAsync();

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
CloudServiceRoleInstanceData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
