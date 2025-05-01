using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.RedisEnterprise.Models;
using Azure.ResourceManager.Resources;
using Azure.ResourceManager.RedisEnterprise;

// Generated from example definition: specification/redisenterprise/resource-manager/Microsoft.Cache/preview/2025-05-01-preview/examples/OperationsStatusGet.json
// this example is just showing the usage of "OperationsStatus_Get" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this SubscriptionResource created on azure
// for more information of creating SubscriptionResource, please refer to the document of SubscriptionResource
string subscriptionId = "e7b5a9d2-6b6a-4d2f-9143-20d9a10f5b8f";
ResourceIdentifier subscriptionResourceId = SubscriptionResource.CreateResourceIdentifier(subscriptionId);
SubscriptionResource subscriptionResource = client.GetSubscriptionResource(subscriptionResourceId);

// invoke the operation
AzureLocation location = new AzureLocation("West US");
string operationId = "testoperationid";
RedisEnterpriseOperationStatus result = await subscriptionResource.GetRedisEnterpriseOperationsStatusAsync(location, operationId);

Console.WriteLine($"Succeeded: {result}");
