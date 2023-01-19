using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.DataShare;
using Azure.ResourceManager.DataShare.Models;

// Generated from example definition: specification/datashare/resource-manager/Microsoft.DataShare/stable/2021-08-01/examples/ShareSubscriptions_Delete.json
// this example is just showing the usage of "ShareSubscriptions_Delete" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ShareSubscriptionResource created on azure
// for more information of creating ShareSubscriptionResource, please refer to the document of ShareSubscriptionResource
string subscriptionId = "12345678-1234-1234-12345678abc";
string resourceGroupName = "SampleResourceGroup";
string accountName = "Account1";
string shareSubscriptionName = "ShareSubscription1";
ResourceIdentifier shareSubscriptionResourceId = ShareSubscriptionResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, accountName, shareSubscriptionName);
ShareSubscriptionResource shareSubscription = client.GetShareSubscriptionResource(shareSubscriptionResourceId);

// invoke the operation
ArmOperation<DataShareOperationResult> lro = await shareSubscription.DeleteAsync(WaitUntil.Completed);
DataShareOperationResult result = lro.Value;

Console.WriteLine($"Succeeded: {result}");
