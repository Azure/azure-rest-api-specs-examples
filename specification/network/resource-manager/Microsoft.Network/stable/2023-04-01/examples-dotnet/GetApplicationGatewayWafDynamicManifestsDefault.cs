using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Network;
using Azure.ResourceManager.Resources;

// Generated from example definition: specification/network/resource-manager/Microsoft.Network/stable/2023-04-01/examples/GetApplicationGatewayWafDynamicManifestsDefault.json
// this example is just showing the usage of "ApplicationGatewayWafDynamicManifestsDefault_Get" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this SubscriptionResource created on azure
// for more information of creating SubscriptionResource, please refer to the document of SubscriptionResource
string subscriptionId = "subid";
ResourceIdentifier subscriptionResourceId = SubscriptionResource.CreateResourceIdentifier(subscriptionId);
SubscriptionResource subscriptionResource = client.GetSubscriptionResource(subscriptionResourceId);

// get the collection of this ApplicationGatewayWafDynamicManifestResource
AzureLocation location = new AzureLocation("westus");
ApplicationGatewayWafDynamicManifestCollection collection = subscriptionResource.GetApplicationGatewayWafDynamicManifests(location);

// invoke the operation
bool result = await collection.ExistsAsync();

Console.WriteLine($"Succeeded: {result}");
