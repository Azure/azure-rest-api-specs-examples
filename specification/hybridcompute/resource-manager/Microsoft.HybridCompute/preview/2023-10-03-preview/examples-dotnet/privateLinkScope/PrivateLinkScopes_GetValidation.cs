using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.HybridCompute;
using Azure.ResourceManager.HybridCompute.Models;
using Azure.ResourceManager.Resources;

// Generated from example definition: specification/hybridcompute/resource-manager/Microsoft.HybridCompute/preview/2023-10-03-preview/examples/privateLinkScope/PrivateLinkScopes_GetValidation.json
// this example is just showing the usage of "PrivateLinkScopes_GetValidationDetails" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this SubscriptionResource created on azure
// for more information of creating SubscriptionResource, please refer to the document of SubscriptionResource
string subscriptionId = "86dc51d3-92ed-4d7e-947a-775ea79b4919";
ResourceIdentifier subscriptionResourceId = SubscriptionResource.CreateResourceIdentifier(subscriptionId);
SubscriptionResource subscriptionResource = client.GetSubscriptionResource(subscriptionResourceId);

// invoke the operation
AzureLocation location = new AzureLocation("wus2");
string privateLinkScopeId = "f5dc51d3-92ed-4d7e-947a-775ea79b4919";
PrivateLinkScopeValidationDetails result = await subscriptionResource.GetValidationDetailsPrivateLinkScopeAsync(location, privateLinkScopeId);

Console.WriteLine($"Succeeded: {result}");
