using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.AlertsManagement;
using Azure.ResourceManager.AlertsManagement.Models;

// Generated from example definition: specification/alertsmanagement/resource-manager/Microsoft.AlertsManagement/preview/2019-05-05-preview/examples/SmartGroups_ChangeState.json
// this example is just showing the usage of "SmartGroups_ChangeState" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this SmartGroupResource created on azure
// for more information of creating SmartGroupResource, please refer to the document of SmartGroupResource
string subscriptionId = "dd91de05-d791-4ceb-b6dc-988682dc7d72";
Guid smartGroupId = Guid.Parse("a808445e-bb38-4751-85c2-1b109ccc1059");
ResourceIdentifier smartGroupResourceId = SmartGroupResource.CreateResourceIdentifier(subscriptionId, smartGroupId);
SmartGroupResource smartGroup = client.GetSmartGroupResource(smartGroupResourceId);

// invoke the operation
ServiceAlertState newState = ServiceAlertState.Acknowledged;
SmartGroupResource result = await smartGroup.ChangeStateAsync(newState);

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
SmartGroupData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
