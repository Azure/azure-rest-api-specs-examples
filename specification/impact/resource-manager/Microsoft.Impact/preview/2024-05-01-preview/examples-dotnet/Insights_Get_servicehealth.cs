using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.ImpactReporting.Models;
using Azure.ResourceManager.ImpactReporting;

// Generated from example definition: 2024-05-01-preview/Insights_Get_servicehealth.json
// this example is just showing the usage of "Insight_Get" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ImpactInsightResource created on azure
// for more information of creating ImpactInsightResource, please refer to the document of ImpactInsightResource
string subscriptionId = "00000000-0000-0000-0000-000000000000";
string workloadImpactName = "impactid";
string insightName = "insightname";
ResourceIdentifier impactInsightResourceId = ImpactInsightResource.CreateResourceIdentifier(subscriptionId, workloadImpactName, insightName);
ImpactInsightResource impactInsight = client.GetImpactInsightResource(impactInsightResourceId);

// invoke the operation
ImpactInsightResource result = await impactInsight.GetAsync();

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
ImpactInsightData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
