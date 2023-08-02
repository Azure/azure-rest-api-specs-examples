using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.DataFactory;
using Azure.ResourceManager.DataFactory.Models;

// Generated from example definition: specification/datafactory/resource-manager/Microsoft.DataFactory/stable/2018-06-01/examples/Triggers_UnsubscribeFromEvents.json
// this example is just showing the usage of "Triggers_UnsubscribeFromEvents" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this DataFactoryTriggerResource created on azure
// for more information of creating DataFactoryTriggerResource, please refer to the document of DataFactoryTriggerResource
string subscriptionId = "12345678-1234-1234-1234-12345678abc";
string resourceGroupName = "exampleResourceGroup";
string factoryName = "exampleFactoryName";
string triggerName = "exampleTrigger";
ResourceIdentifier dataFactoryTriggerResourceId = DataFactoryTriggerResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, factoryName, triggerName);
DataFactoryTriggerResource dataFactoryTrigger = client.GetDataFactoryTriggerResource(dataFactoryTriggerResourceId);

// invoke the operation
ArmOperation<DataFactoryTriggerSubscriptionOperationResult> lro = await dataFactoryTrigger.UnsubscribeFromEventsAsync(WaitUntil.Completed);
DataFactoryTriggerSubscriptionOperationResult result = lro.Value;

Console.WriteLine($"Succeeded: {result}");
