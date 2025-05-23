using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.MixedReality.Models;
using Azure.ResourceManager.Resources;
using Azure.ResourceManager.MixedReality;

// Generated from example definition: specification/mixedreality/resource-manager/Microsoft.MixedReality/stable/2021-01-01/examples/proxy/CheckNameAvailabilityForLocalUniqueness.json
// this example is just showing the usage of "CheckNameAvailabilityLocal" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this SubscriptionResource created on azure
// for more information of creating SubscriptionResource, please refer to the document of SubscriptionResource
string subscriptionId = "xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx";
ResourceIdentifier subscriptionResourceId = SubscriptionResource.CreateResourceIdentifier(subscriptionId);
SubscriptionResource subscriptionResource = client.GetSubscriptionResource(subscriptionResourceId);

// invoke the operation
AzureLocation location = new AzureLocation("eastus2euap");
MixedRealityNameAvailabilityContent content = new MixedRealityNameAvailabilityContent("MyAccount", "Microsoft.MixedReality/spatialAnchorsAccounts");
MixedRealityNameAvailabilityResult result = await subscriptionResource.CheckMixedRealityNameAvailabilityAsync(location, content);

Console.WriteLine($"Succeeded: {result}");
