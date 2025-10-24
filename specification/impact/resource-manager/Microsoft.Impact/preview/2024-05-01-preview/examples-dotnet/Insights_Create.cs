using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.ImpactReporting.Models;
using Azure.ResourceManager.ImpactReporting;

// Generated from example definition: 2024-05-01-preview/Insights_Create.json
// this example is just showing the usage of "Insight_Create" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this WorkloadImpactResource created on azure
// for more information of creating WorkloadImpactResource, please refer to the document of WorkloadImpactResource
string subscriptionId = "00000000-0000-0000-0000-000000000000";
string workloadImpactName = "impactid22";
ResourceIdentifier workloadImpactResourceId = WorkloadImpactResource.CreateResourceIdentifier(subscriptionId, workloadImpactName);
WorkloadImpactResource workloadImpact = client.GetWorkloadImpactResource(workloadImpactResourceId);

// get the collection of this ImpactInsightResource
ImpactInsightCollection collection = workloadImpact.GetImpactInsights();

// invoke the operation
string insightName = "insightId12";
ImpactInsightData data = new ImpactInsightData
{
    Properties = new ImpactInsightProperties("repair", new ImpactInsightContent("Impact Has been correlated to an outage", "At 2018-11-08T00:00:00Z UTC, your services dependent on these resources <link href=”…”>VM1</link> may have experienced an issue. <br/><div>We have identified an outage that affected these resources(s). You can look at outage information on <link href=\"https:// portal.azure.com/#view/Microsoft_Azure_Health/AzureHealthBrowseBlade/~/serviceIssues/trackingId/NL2W-VCZ\">NL2W-VCZ</link> link.<div>"), "00000000-0000-0000-0000-000000000000", new ImpactDetails(new ResourceIdentifier("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resource-rg/providers/Microsoft.Sql/sqlserver/dbservername"), DateTimeOffset.Parse("2023-06-15T01:00:00.009223Z"), new ResourceIdentifier("/subscriptions/00000000-0000-0000-0000-000000000000/providers/microsoft.Impact/workloadImpacts/impactid22")))
    {
        Status = "resolved",
        EventOn = DateTimeOffset.Parse("2023-06-15T04:00:00.009223Z"),
    },
};
ArmOperation<ImpactInsightResource> lro = await collection.CreateOrUpdateAsync(WaitUntil.Completed, insightName, data);
ImpactInsightResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
ImpactInsightData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
