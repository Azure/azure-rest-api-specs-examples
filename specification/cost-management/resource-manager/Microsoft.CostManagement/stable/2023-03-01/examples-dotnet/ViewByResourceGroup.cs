using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.CostManagement;
using Azure.ResourceManager.CostManagement.Models;

// Generated from example definition: specification/cost-management/resource-manager/Microsoft.CostManagement/stable/2023-03-01/examples/ViewByResourceGroup.json
// this example is just showing the usage of "Views_GetByScope" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ArmResource created on azure
// for more information of creating ArmResource, please refer to the document of ArmResource

// get the collection of this CostManagementViewsResource
string scope = "subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/MYDEVTESTRG";
ResourceIdentifier scopeId = new ResourceIdentifier(string.Format("/{0}", scope));
CostManagementViewsCollection collection = client.GetAllCostManagementViews(scopeId);

// invoke the operation
string viewName = "swaggerExample";
bool result = await collection.ExistsAsync(viewName);

Console.WriteLine($"Succeeded: {result}");
