using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.ServiceBus;

// Generated from example definition: specification/servicebus/resource-manager/Microsoft.ServiceBus/preview/2022-10-01-preview/examples/disasterRecoveryConfigs/SBAliasCreate.json
// this example is just showing the usage of "DisasterRecoveryConfigs_CreateOrUpdate" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ServiceBusDisasterRecoveryResource created on azure
// for more information of creating ServiceBusDisasterRecoveryResource, please refer to the document of ServiceBusDisasterRecoveryResource
string subscriptionId = "5f750a97-50d9-4e36-8081-c9ee4c0210d4";
string resourceGroupName = "ardsouzatestRG";
string namespaceName = "sdk-Namespace-8860";
string @alias = "sdk-Namespace-8860";
ResourceIdentifier serviceBusDisasterRecoveryResourceId = ServiceBusDisasterRecoveryResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, namespaceName, @alias);
ServiceBusDisasterRecoveryResource serviceBusDisasterRecovery = client.GetServiceBusDisasterRecoveryResource(serviceBusDisasterRecoveryResourceId);

// invoke the operation
ServiceBusDisasterRecoveryData data = new ServiceBusDisasterRecoveryData()
{
    PartnerNamespace = "sdk-Namespace-37",
    AlternateName = "alternameforAlias-Namespace-8860",
};
ArmOperation<ServiceBusDisasterRecoveryResource> lro = await serviceBusDisasterRecovery.UpdateAsync(WaitUntil.Completed, data);
ServiceBusDisasterRecoveryResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
ServiceBusDisasterRecoveryData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
