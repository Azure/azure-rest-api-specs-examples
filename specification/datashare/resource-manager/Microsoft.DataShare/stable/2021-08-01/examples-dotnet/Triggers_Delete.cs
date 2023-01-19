using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.DataShare;
using Azure.ResourceManager.DataShare.Models;

// Generated from example definition: specification/datashare/resource-manager/Microsoft.DataShare/stable/2021-08-01/examples/Triggers_Delete.json
// this example is just showing the usage of "Triggers_Delete" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this DataShareTriggerResource created on azure
// for more information of creating DataShareTriggerResource, please refer to the document of DataShareTriggerResource
string subscriptionId = "433a8dfd-e5d5-4e77-ad86-90acdc75eb1a";
string resourceGroupName = "SampleResourceGroup";
string accountName = "Account1";
string shareSubscriptionName = "ShareSubscription1";
string triggerName = "Trigger1";
ResourceIdentifier dataShareTriggerResourceId = DataShareTriggerResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, accountName, shareSubscriptionName, triggerName);
DataShareTriggerResource dataShareTrigger = client.GetDataShareTriggerResource(dataShareTriggerResourceId);

// invoke the operation
ArmOperation<DataShareOperationResult> lro = await dataShareTrigger.DeleteAsync(WaitUntil.Completed);
DataShareOperationResult result = lro.Value;

Console.WriteLine($"Succeeded: {result}");
