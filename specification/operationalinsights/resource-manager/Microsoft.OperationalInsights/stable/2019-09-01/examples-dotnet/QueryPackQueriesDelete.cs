using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.OperationalInsights.Models;
using Azure.ResourceManager.OperationalInsights;

// Generated from example definition: specification/operationalinsights/resource-manager/Microsoft.OperationalInsights/stable/2019-09-01/examples/QueryPackQueriesDelete.json
// this example is just showing the usage of "Queries_Delete" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this LogAnalyticsQueryResource created on azure
// for more information of creating LogAnalyticsQueryResource, please refer to the document of LogAnalyticsQueryResource
string subscriptionId = "86dc51d3-92ed-4d7e-947a-775ea79b4918";
string resourceGroupName = "my-resource-group";
string queryPackName = "my-querypack";
string id = "a449f8af-8e64-4b3a-9b16-5a7165ff98c4";
ResourceIdentifier logAnalyticsQueryResourceId = LogAnalyticsQueryResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, queryPackName, id);
LogAnalyticsQueryResource logAnalyticsQuery = client.GetLogAnalyticsQueryResource(logAnalyticsQueryResourceId);

// invoke the operation
await logAnalyticsQuery.DeleteAsync(WaitUntil.Completed);

Console.WriteLine($"Succeeded");
