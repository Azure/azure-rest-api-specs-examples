using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.CustomerInsights;
using Azure.ResourceManager.CustomerInsights.Models;

// Generated from example definition: specification/customer-insights/resource-manager/Microsoft.CustomerInsights/stable/2017-04-26/examples/PredictionsCreateOrUpdate.json
// this example is just showing the usage of "Predictions_CreateOrUpdate" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this HubResource created on azure
// for more information of creating HubResource, please refer to the document of HubResource
string subscriptionId = "c909e979-ef71-4def-a970-bc7c154db8c5";
string resourceGroupName = "TestHubRG";
string hubName = "sdkTestHub";
ResourceIdentifier hubResourceId = HubResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, hubName);
HubResource hub = client.GetHubResource(hubResourceId);

// get the collection of this PredictionResourceFormatResource
PredictionResourceFormatCollection collection = hub.GetPredictionResourceFormats();

// invoke the operation
string predictionName = "sdktest";
PredictionResourceFormatData data = new PredictionResourceFormatData()
{
    Description =
    {
    ["en-us"] = "sdktest",
    },
    DisplayName =
    {
    ["en-us"] = "sdktest",
    },
    InvolvedInteractionTypes =
    {
    },
    InvolvedKpiTypes =
    {
    },
    InvolvedRelationships =
    {
    },
    NegativeOutcomeExpression = "Customers.FirstName = 'Mike'",
    PositiveOutcomeExpression = "Customers.FirstName = 'David'",
    PrimaryProfileType = "Customers",
    PredictionName = "sdktest",
    ScopeExpression = "*",
    AutoAnalyze = true,
    Mappings = new PredictionMappings("sdktest_Score", "sdktest_Grade", "sdktest_Reason"),
    ScoreLabel = "score label",
    Grades =
    {
    },
};
ArmOperation<PredictionResourceFormatResource> lro = await collection.CreateOrUpdateAsync(WaitUntil.Completed, predictionName, data);
PredictionResourceFormatResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
PredictionResourceFormatData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
