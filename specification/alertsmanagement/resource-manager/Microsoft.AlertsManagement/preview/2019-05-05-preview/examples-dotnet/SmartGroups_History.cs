using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.AlertsManagement;
using Azure.ResourceManager.AlertsManagement.Models;

// Generated from example definition: specification/alertsmanagement/resource-manager/Microsoft.AlertsManagement/preview/2019-05-05-preview/examples/SmartGroups_History.json
// this example is just showing the usage of "SmartGroups_GetHistory" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this SmartGroupResource created on azure
// for more information of creating SmartGroupResource, please refer to the document of SmartGroupResource
string subscriptionId = "9e261de7-c804-4b9d-9ebf-6f50fe350a9a";
Guid smartGroupId = Guid.Parse("a808445e-bb38-4751-85c2-1b109ccc1059");
ResourceIdentifier smartGroupResourceId = SmartGroupResource.CreateResourceIdentifier(subscriptionId, smartGroupId);
SmartGroupResource smartGroup = client.GetSmartGroupResource(smartGroupResourceId);

// invoke the operation
SmartGroupModification result = await smartGroup.GetHistoryAsync();

Console.WriteLine($"Succeeded: {result}");
