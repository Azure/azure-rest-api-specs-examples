using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.AlertsManagement;
using Azure.ResourceManager.AlertsManagement.Models;
using Azure.ResourceManager.Resources;

// Generated from example definition: specification/alertsmanagement/resource-manager/Microsoft.AlertsManagement/stable/2021-08-08/examples/AlertProcessingRules_List_ResourceGroup.json
// this example is just showing the usage of "AlertProcessingRules_ListByResourceGroup" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ResourceGroupResource created on azure
// for more information of creating ResourceGroupResource, please refer to the document of ResourceGroupResource
string subscriptionId = "1e3ff1c0-771a-4119-a03b-be82a51e232d";
string resourceGroupName = "alertscorrelationrg";
ResourceIdentifier resourceGroupResourceId = ResourceGroupResource.CreateResourceIdentifier(subscriptionId, resourceGroupName);
ResourceGroupResource resourceGroupResource = client.GetResourceGroupResource(resourceGroupResourceId);

// get the collection of this AlertProcessingRuleResource
AlertProcessingRuleCollection collection = resourceGroupResource.GetAlertProcessingRules();

// invoke the operation and iterate over the result
await foreach (AlertProcessingRuleResource item in collection.GetAllAsync())
{
    // the variable item is a resource, you could call other operations on this instance as well
    // but just for demo, we get its data from this resource instance
    AlertProcessingRuleData resourceData = item.Data;
    // for demo we just print out the id
    Console.WriteLine($"Succeeded on id: {resourceData.Id}");
}

Console.WriteLine($"Succeeded");
