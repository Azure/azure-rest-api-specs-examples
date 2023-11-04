using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.MySql.FlexibleServers;
using Azure.ResourceManager.MySql.FlexibleServers.Models;
using Azure.ResourceManager.Resources;

// Generated from example definition: specification/mysql/resource-manager/Microsoft.DBforMySQL/ServiceOperations/preview/2023-06-01-preview/examples/OperationResults_Get.json
// this example is just showing the usage of "OperationResults_Get" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this SubscriptionResource created on azure
// for more information of creating SubscriptionResource, please refer to the document of SubscriptionResource
string subscriptionId = "11111111-2222-3333-4444-555555555555";
ResourceIdentifier subscriptionResourceId = SubscriptionResource.CreateResourceIdentifier(subscriptionId);
SubscriptionResource subscriptionResource = client.GetSubscriptionResource(subscriptionResourceId);

// invoke the operation
AzureLocation locationName = new AzureLocation("westus");
string operationId = "resource-provisioning-id-farmBeatsResourceName";
OperationStatusExtendedResult result = await subscriptionResource.GetOperationResultAsync(locationName, operationId);

Console.WriteLine($"Succeeded: {result}");
