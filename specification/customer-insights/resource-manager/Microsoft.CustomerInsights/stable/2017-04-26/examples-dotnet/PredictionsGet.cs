using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.CustomerInsights;
using Azure.ResourceManager.CustomerInsights.Models;

// Generated from example definition: specification/customer-insights/resource-manager/Microsoft.CustomerInsights/stable/2017-04-26/examples/PredictionsGet.json
// this example is just showing the usage of "Predictions_Get" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this PredictionResourceFormatResource created on azure
// for more information of creating PredictionResourceFormatResource, please refer to the document of PredictionResourceFormatResource
string subscriptionId = "c909e979-ef71-4def-a970-bc7c154db8c5";
string resourceGroupName = "TestHubRG";
string hubName = "sdkTestHub";
string predictionName = "sdktest";
ResourceIdentifier predictionResourceFormatResourceId = PredictionResourceFormatResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, hubName, predictionName);
PredictionResourceFormatResource predictionResourceFormat = client.GetPredictionResourceFormatResource(predictionResourceFormatResourceId);

// invoke the operation
PredictionResourceFormatResource result = await predictionResourceFormat.GetAsync();

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
PredictionResourceFormatData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
