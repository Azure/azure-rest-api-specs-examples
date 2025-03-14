using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.AlertsManagement.Models;
using Azure.ResourceManager.AlertsManagement;

// Generated from example definition: specification/alertsmanagement/resource-manager/Microsoft.AlertsManagement/preview/2019-05-05-preview/examples/SmartGroups_GetById.json
// this example is just showing the usage of "SmartGroups_GetById" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this SmartGroupResource created on azure
// for more information of creating SmartGroupResource, please refer to the document of SmartGroupResource
string subscriptionId = "9e261de7-c804-4b9d-9ebf-6f50fe350a9a";
Guid smartGroupId = Guid.Parse("603675da-9851-4b26-854a-49fc53d32715");
ResourceIdentifier smartGroupResourceId = SmartGroupResource.CreateResourceIdentifier(subscriptionId, smartGroupId);
SmartGroupResource smartGroup = client.GetSmartGroupResource(smartGroupResourceId);

// invoke the operation
SmartGroupResource result = await smartGroup.GetAsync();

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
SmartGroupData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
