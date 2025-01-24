using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.IotOperations.Models;
using Azure.ResourceManager.IotOperations;

// Generated from example definition: 2024-11-01/Broker_CreateOrUpdate_Minimal.json
// this example is just showing the usage of "BrokerResource_CreateOrUpdate" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this IotOperationsBrokerResource created on azure
// for more information of creating IotOperationsBrokerResource, please refer to the document of IotOperationsBrokerResource
string subscriptionId = "F8C729F9-DF9C-4743-848F-96EE433D8E53";
string resourceGroupName = "rgiotoperations";
string instanceName = "resource-name123";
string brokerName = "resource-name123";
ResourceIdentifier iotOperationsBrokerResourceId = IotOperationsBrokerResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, instanceName, brokerName);
IotOperationsBrokerResource iotOperationsBroker = client.GetIotOperationsBrokerResource(iotOperationsBrokerResourceId);

// invoke the operation
IotOperationsBrokerData data = new IotOperationsBrokerData(new IotOperationsExtendedLocation("qmbrfwcpwwhggszhrdjv", IotOperationsExtendedLocationType.CustomLocation))
{
    Properties = new IotOperationsBrokerProperties
    {
        MemoryProfile = BrokerMemoryProfile.Tiny,
    },
};
ArmOperation<IotOperationsBrokerResource> lro = await iotOperationsBroker.UpdateAsync(WaitUntil.Completed, data);
IotOperationsBrokerResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
IotOperationsBrokerData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
