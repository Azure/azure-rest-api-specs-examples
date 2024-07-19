using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.OperationalInsights.Models;
using Azure.ResourceManager.Resources;
using Azure.ResourceManager.OperationalInsights;

// Generated from example definition: specification/operationalinsights/resource-manager/Microsoft.OperationalInsights/stable/2019-09-01/examples/QueryPacksDelete.json
// this example is just showing the usage of "QueryPacks_Delete" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this LogAnalyticsQueryPackResource created on azure
// for more information of creating LogAnalyticsQueryPackResource, please refer to the document of LogAnalyticsQueryPackResource
string subscriptionId = "86dc51d3-92ed-4d7e-947a-775ea79b4919";
string resourceGroupName = "my-resource-group";
string queryPackName = "my-querypack";
ResourceIdentifier logAnalyticsQueryPackResourceId = LogAnalyticsQueryPackResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, queryPackName);
LogAnalyticsQueryPackResource logAnalyticsQueryPack = client.GetLogAnalyticsQueryPackResource(logAnalyticsQueryPackResourceId);

// invoke the operation
await logAnalyticsQueryPack.DeleteAsync(WaitUntil.Completed);

Console.WriteLine($"Succeeded");
