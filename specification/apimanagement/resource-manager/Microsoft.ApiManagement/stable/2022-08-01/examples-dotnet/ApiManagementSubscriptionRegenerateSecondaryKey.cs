using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.ApiManagement.Models;
using Azure.ResourceManager.ApiManagement;

// Generated from example definition: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2022-08-01/examples/ApiManagementSubscriptionRegenerateSecondaryKey.json
// this example is just showing the usage of "Subscription_RegenerateSecondaryKey" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ApiManagementSubscriptionResource created on azure
// for more information of creating ApiManagementSubscriptionResource, please refer to the document of ApiManagementSubscriptionResource
string subscriptionId = "subid";
string resourceGroupName = "rg1";
string serviceName = "apimService1";
string sid = "testsub";
ResourceIdentifier apiManagementSubscriptionResourceId = ApiManagementSubscriptionResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, serviceName, sid);
ApiManagementSubscriptionResource apiManagementSubscription = client.GetApiManagementSubscriptionResource(apiManagementSubscriptionResourceId);

// invoke the operation
await apiManagementSubscription.RegenerateSecondaryKeyAsync();

Console.WriteLine($"Succeeded");
