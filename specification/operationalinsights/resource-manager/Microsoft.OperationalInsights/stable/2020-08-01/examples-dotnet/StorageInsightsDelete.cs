using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.OperationalInsights.Models;
using Azure.ResourceManager.OperationalInsights;

// Generated from example definition: specification/operationalinsights/resource-manager/Microsoft.OperationalInsights/stable/2020-08-01/examples/StorageInsightsDelete.json
// this example is just showing the usage of "StorageInsightConfigs_Delete" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this StorageInsightResource created on azure
// for more information of creating StorageInsightResource, please refer to the document of StorageInsightResource
string subscriptionId = "00000000-0000-0000-0000-00000000000";
string resourceGroupName = "OIAutoRest5123";
string workspaceName = "aztest5048";
string storageInsightName = "AzTestSI1110";
ResourceIdentifier storageInsightResourceId = StorageInsightResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, workspaceName, storageInsightName);
StorageInsightResource storageInsight = client.GetStorageInsightResource(storageInsightResourceId);

// invoke the operation
await storageInsight.DeleteAsync(WaitUntil.Completed);

Console.WriteLine($"Succeeded");
