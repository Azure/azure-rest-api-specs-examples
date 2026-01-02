using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.IotOperations.Models;
using Azure.ResourceManager.IotOperations;

// Generated from example definition: 2025-07-01-preview/Dataflow_ListByProfileResource_MaximumSet_Gen.json
// this example is just showing the usage of "DataflowResource_ListByResourceGroup" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this IotOperationsDataflowProfileResource created on azure
// for more information of creating IotOperationsDataflowProfileResource, please refer to the document of IotOperationsDataflowProfileResource
string subscriptionId = "F8C729F9-DF9C-4743-848F-96EE433D8E53";
string resourceGroupName = "rgiotoperations";
string instanceName = "resource-name123";
string dataflowProfileName = "resource-name123";
ResourceIdentifier iotOperationsDataflowProfileResourceId = IotOperationsDataflowProfileResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, instanceName, dataflowProfileName);
IotOperationsDataflowProfileResource iotOperationsDataflowProfile = client.GetIotOperationsDataflowProfileResource(iotOperationsDataflowProfileResourceId);

// get the collection of this IotOperationsDataflowResource
IotOperationsDataflowCollection collection = iotOperationsDataflowProfile.GetIotOperationsDataflows();

// invoke the operation and iterate over the result
await foreach (IotOperationsDataflowResource item in collection.GetAllAsync())
{
    // the variable item is a resource, you could call other operations on this instance as well
    // but just for demo, we get its data from this resource instance
    IotOperationsDataflowData resourceData = item.Data;
    // for demo we just print out the id
    Console.WriteLine($"Succeeded on id: {resourceData.Id}");
}

Console.WriteLine("Succeeded");
