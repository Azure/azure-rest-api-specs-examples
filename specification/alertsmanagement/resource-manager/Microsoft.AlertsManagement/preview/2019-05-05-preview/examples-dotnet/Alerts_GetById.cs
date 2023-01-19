using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.AlertsManagement;
using Azure.ResourceManager.AlertsManagement.Models;
using Azure.ResourceManager.Resources;

// Generated from example definition: specification/alertsmanagement/resource-manager/Microsoft.AlertsManagement/preview/2019-05-05-preview/examples/Alerts_GetById.json
// this example is just showing the usage of "Alerts_GetById" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ServiceAlertResource created on azure
// for more information of creating ServiceAlertResource, please refer to the document of ServiceAlertResource
string subscriptionId = "9e261de7-c804-4b9d-9ebf-6f50fe350a9a";
Guid alertId = Guid.Parse("66114d64-d9d9-478b-95c9-b789d6502100");
ResourceIdentifier serviceAlertResourceId = ServiceAlertResource.CreateResourceIdentifier(subscriptionId, alertId);
ServiceAlertResource serviceAlert = client.GetServiceAlertResource(serviceAlertResourceId);

// invoke the operation
ServiceAlertResource result = await serviceAlert.GetAsync();

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
ServiceAlertData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
