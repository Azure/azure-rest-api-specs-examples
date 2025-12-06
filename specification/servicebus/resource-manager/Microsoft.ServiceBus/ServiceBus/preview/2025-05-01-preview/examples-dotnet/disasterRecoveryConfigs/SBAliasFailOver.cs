using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.ServiceBus;

// Generated from example definition: specification/servicebus/resource-manager/Microsoft.ServiceBus/ServiceBus/preview/2025-05-01-preview/examples/disasterRecoveryConfigs/SBAliasFailOver.json
// this example is just showing the usage of "DisasterRecoveryConfigs_FailOver" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ServiceBusDisasterRecoveryResource created on azure
// for more information of creating ServiceBusDisasterRecoveryResource, please refer to the document of ServiceBusDisasterRecoveryResource
string subscriptionId = "5f750a97-50d9-4e36-8081-c9ee4c0210d4";
string resourceGroupName = "ardsouzatestRG";
string namespaceName = "sdk-Namespace-8860";
string @alias = "sdk-DisasterRecovery-3814";
ResourceIdentifier serviceBusDisasterRecoveryResourceId = ServiceBusDisasterRecoveryResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, namespaceName, @alias);
ServiceBusDisasterRecoveryResource serviceBusDisasterRecovery = client.GetServiceBusDisasterRecoveryResource(serviceBusDisasterRecoveryResourceId);

// invoke the operation
await serviceBusDisasterRecovery.FailOverAsync();

Console.WriteLine("Succeeded");
