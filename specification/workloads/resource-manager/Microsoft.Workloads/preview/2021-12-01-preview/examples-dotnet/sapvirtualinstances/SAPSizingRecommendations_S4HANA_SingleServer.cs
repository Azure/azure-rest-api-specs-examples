using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Resources;
using Azure.ResourceManager.Workloads;
using Azure.ResourceManager.Workloads.Models;

// Generated from example definition: specification/workloads/resource-manager/Microsoft.Workloads/preview/2021-12-01-preview/examples/sapvirtualinstances/SAPSizingRecommendations_S4HANA_SingleServer.json
// this example is just showing the usage of "SAPSizingRecommendations" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this SubscriptionResource created on azure
// for more information of creating SubscriptionResource, please refer to the document of SubscriptionResource
string subscriptionId = "8e17e36c-42e9-4cd5-a078-7b44883414e0";
ResourceIdentifier subscriptionResourceId = SubscriptionResource.CreateResourceIdentifier(subscriptionId);
SubscriptionResource subscriptionResource = client.GetSubscriptionResource(subscriptionResourceId);

// invoke the operation
AzureLocation location = new AzureLocation("centralus");
SapSizingRecommendationResult result = await subscriptionResource.SapSizingRecommendationsAsync(location);

Console.WriteLine($"Succeeded: {result}");
