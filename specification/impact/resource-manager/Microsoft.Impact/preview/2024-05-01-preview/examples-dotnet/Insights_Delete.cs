using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.ImpactReporting.Models;
using Azure.ResourceManager.ImpactReporting;

// Generated from example definition: 2024-05-01-preview/Insights_Delete.json
// this example is just showing the usage of "Insight_Delete" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ImpactInsightResource created on azure
// for more information of creating ImpactInsightResource, please refer to the document of ImpactInsightResource
string subscriptionId = "00000000-0000-0000-0000-000000000000";
string workloadImpactName = "impactid22";
string insightName = "insightId12";
ResourceIdentifier impactInsightResourceId = ImpactInsightResource.CreateResourceIdentifier(subscriptionId, workloadImpactName, insightName);
ImpactInsightResource impactInsight = client.GetImpactInsightResource(impactInsightResourceId);

// invoke the operation
await impactInsight.DeleteAsync(WaitUntil.Completed);

Console.WriteLine("Succeeded");
