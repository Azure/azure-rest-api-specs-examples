using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.CostManagement.Models;
using Azure.ResourceManager.CostManagement;

// Generated from example definition: specification/cost-management/resource-manager/Microsoft.CostManagement/stable/2023-03-01/examples/ViewListByResourceGroup.json
// this example is just showing the usage of "Views_ListByScope" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// get the collection of this CostManagementViewsResource
string scope = "subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/MYDEVTESTRG";
CostManagementViewsCollection collection = client.GetAllCostManagementViews(new ResourceIdentifier(scope));

// invoke the operation and iterate over the result
await foreach (CostManagementViewsResource item in collection.GetAllAsync())
{
    // the variable item is a resource, you could call other operations on this instance as well
    // but just for demo, we get its data from this resource instance
    CostManagementViewData resourceData = item.Data;
    // for demo we just print out the id
    Console.WriteLine($"Succeeded on id: {resourceData.Id}");
}

Console.WriteLine("Succeeded");
