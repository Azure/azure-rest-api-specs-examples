using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.ApiManagement.Models;
using Azure.ResourceManager.ApiManagement;

// Generated from example definition: specification/apimanagement/resource-manager/Microsoft.ApiManagement/preview/2023-03-01-preview/examples/ApiManagementSubscriptionListSecrets.json
// this example is just showing the usage of "Subscription_ListSecrets" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ApiManagementSubscriptionResource created on azure
// for more information of creating ApiManagementSubscriptionResource, please refer to the document of ApiManagementSubscriptionResource
string subscriptionId = "00000000-0000-0000-0000-000000000000";
string resourceGroupName = "rg1";
string serviceName = "apimService1";
string sid = "5931a769d8d14f0ad8ce13b8";
ResourceIdentifier apiManagementSubscriptionResourceId = ApiManagementSubscriptionResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, serviceName, sid);
ApiManagementSubscriptionResource apiManagementSubscription = client.GetApiManagementSubscriptionResource(apiManagementSubscriptionResourceId);

// invoke the operation
SubscriptionKeysContract result = await apiManagementSubscription.GetSecretsAsync();

Console.WriteLine($"Succeeded: {result}");
