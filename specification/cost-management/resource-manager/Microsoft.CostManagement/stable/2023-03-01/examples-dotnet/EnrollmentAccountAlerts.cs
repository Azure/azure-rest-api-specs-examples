using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.CostManagement;

// Generated from example definition: specification/cost-management/resource-manager/Microsoft.CostManagement/stable/2023-03-01/examples/EnrollmentAccountAlerts.json
// this example is just showing the usage of "Alerts_List" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ArmResource created on azure
// for more information of creating ArmResource, please refer to the document of ArmResource

// get the collection of this CostManagementAlertResource
string scope = "providers/Microsoft.Billing/billingAccounts/12345:6789/enrollmentAccounts/456";
ResourceIdentifier scopeId = new ResourceIdentifier(string.Format("/{0}", scope));
CostManagementAlertCollection collection = client.GetCostManagementAlerts(scopeId);

// invoke the operation and iterate over the result
await foreach (CostManagementAlertResource item in collection.GetAllAsync())
{
    // the variable item is a resource, you could call other operations on this instance as well
    // but just for demo, we get its data from this resource instance
    CostManagementAlertData resourceData = item.Data;
    // for demo we just print out the id
    Console.WriteLine($"Succeeded on id: {resourceData.Id}");
}

Console.WriteLine($"Succeeded");
