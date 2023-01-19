using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.DataMigration;
using Azure.ResourceManager.DataMigration.Models;

// Generated from example definition: specification/datamigration/resource-manager/Microsoft.DataMigration/preview/2022-03-30-preview/examples/Tasks_Get.json
// this example is just showing the usage of "Tasks_Get" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ServiceProjectTaskResource created on azure
// for more information of creating ServiceProjectTaskResource, please refer to the document of ServiceProjectTaskResource
string subscriptionId = "fc04246f-04c5-437e-ac5e-206a19e7193f";
string groupName = "DmsSdkRg";
string serviceName = "DmsSdkService";
string projectName = "DmsSdkProject";
string taskName = "DmsSdkTask";
ResourceIdentifier serviceProjectTaskResourceId = ServiceProjectTaskResource.CreateResourceIdentifier(subscriptionId, groupName, serviceName, projectName, taskName);
ServiceProjectTaskResource serviceProjectTask = client.GetServiceProjectTaskResource(serviceProjectTaskResourceId);

// invoke the operation
ServiceProjectTaskResource result = await serviceProjectTask.GetAsync();

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
ProjectTaskData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
