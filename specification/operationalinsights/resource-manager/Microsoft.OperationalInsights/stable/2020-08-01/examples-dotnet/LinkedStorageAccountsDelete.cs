using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.OperationalInsights.Models;
using Azure.ResourceManager.OperationalInsights;

// Generated from example definition: specification/operationalinsights/resource-manager/Microsoft.OperationalInsights/stable/2020-08-01/examples/LinkedStorageAccountsDelete.json
// this example is just showing the usage of "LinkedStorageAccounts_Delete" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this OperationalInsightsLinkedStorageAccountsResource created on azure
// for more information of creating OperationalInsightsLinkedStorageAccountsResource, please refer to the document of OperationalInsightsLinkedStorageAccountsResource
string subscriptionId = "00000000-0000-0000-0000-00000000000";
string resourceGroupName = "mms-eus";
string workspaceName = "testLinkStorageAccountsWS";
OperationalInsightsDataSourceType dataSourceType = OperationalInsightsDataSourceType.CustomLogs;
ResourceIdentifier operationalInsightsLinkedStorageAccountsResourceId = OperationalInsightsLinkedStorageAccountsResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, workspaceName, dataSourceType);
OperationalInsightsLinkedStorageAccountsResource operationalInsightsLinkedStorageAccounts = client.GetOperationalInsightsLinkedStorageAccountsResource(operationalInsightsLinkedStorageAccountsResourceId);

// invoke the operation
await operationalInsightsLinkedStorageAccounts.DeleteAsync(WaitUntil.Completed);

Console.WriteLine($"Succeeded");
