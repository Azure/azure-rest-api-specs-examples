using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.OperationalInsights;
using Azure.ResourceManager.OperationalInsights.Models;

// Generated from example definition: specification/operationalinsights/resource-manager/Microsoft.OperationalInsights/stable/2019-09-01/examples/QueryPackQueriesGet.json
// this example is just showing the usage of "Queries_Get" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this LogAnalyticsQueryPackResource created on azure
// for more information of creating LogAnalyticsQueryPackResource, please refer to the document of LogAnalyticsQueryPackResource
string subscriptionId = "86dc51d3-92ed-4d7e-947a-775ea79b4918";
string resourceGroupName = "my-resource-group";
string queryPackName = "my-querypack";
ResourceIdentifier logAnalyticsQueryPackResourceId = LogAnalyticsQueryPackResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, queryPackName);
LogAnalyticsQueryPackResource logAnalyticsQueryPack = client.GetLogAnalyticsQueryPackResource(logAnalyticsQueryPackResourceId);

// get the collection of this LogAnalyticsQueryResource
LogAnalyticsQueryCollection collection = logAnalyticsQueryPack.GetLogAnalyticsQueries();

// invoke the operation
string id = "a449f8af-8e64-4b3a-9b16-5a7165ff98c4";
bool result = await collection.ExistsAsync(id);

Console.WriteLine($"Succeeded: {result}");
