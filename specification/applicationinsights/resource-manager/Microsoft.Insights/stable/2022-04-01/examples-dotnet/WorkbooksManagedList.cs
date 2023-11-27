using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.ApplicationInsights;
using Azure.ResourceManager.ApplicationInsights.Models;
using Azure.ResourceManager.Models;
using Azure.ResourceManager.Resources;

// Generated from example definition: specification/applicationinsights/resource-manager/Microsoft.Insights/stable/2022-04-01/examples/WorkbooksManagedList.json
// this example is just showing the usage of "Workbooks_ListByResourceGroup" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ResourceGroupResource created on azure
// for more information of creating ResourceGroupResource, please refer to the document of ResourceGroupResource
string subscriptionId = "6b643656-33eb-422f-aee8-3ac119r124af";
string resourceGroupName = "my-resource-group";
ResourceIdentifier resourceGroupResourceId = ResourceGroupResource.CreateResourceIdentifier(subscriptionId, resourceGroupName);
ResourceGroupResource resourceGroupResource = client.GetResourceGroupResource(resourceGroupResourceId);

// get the collection of this WorkbookResource
WorkbookCollection collection = resourceGroupResource.GetWorkbooks();

// invoke the operation and iterate over the result
CategoryType category = CategoryType.Workbook;
string sourceId = "/subscriptions/6b643656-33eb-422f-aee8-3ac119r124af/resourceGroups/my-resource-group/providers/Microsoft.Web/sites/MyApp";
await foreach (WorkbookResource item in collection.GetAllAsync(category, sourceId: sourceId))
{
    // the variable item is a resource, you could call other operations on this instance as well
    // but just for demo, we get its data from this resource instance
    WorkbookData resourceData = item.Data;
    // for demo we just print out the id
    Console.WriteLine($"Succeeded on id: {resourceData.Id}");
}

Console.WriteLine($"Succeeded");
