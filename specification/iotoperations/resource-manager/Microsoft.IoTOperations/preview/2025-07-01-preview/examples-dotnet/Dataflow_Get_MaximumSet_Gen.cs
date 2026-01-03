using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.IotOperations.Models;
using Azure.ResourceManager.IotOperations;

// Generated from example definition: 2025-07-01-preview/Dataflow_Get_MaximumSet_Gen.json
// this example is just showing the usage of "DataflowResource_Get" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this IotOperationsDataflowResource created on azure
// for more information of creating IotOperationsDataflowResource, please refer to the document of IotOperationsDataflowResource
string subscriptionId = "F8C729F9-DF9C-4743-848F-96EE433D8E53";
string resourceGroupName = "rgiotoperations";
string instanceName = "resource-name123";
string dataflowProfileName = "resource-name123";
string dataflowName = "resource-name123";
ResourceIdentifier iotOperationsDataflowResourceId = IotOperationsDataflowResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, instanceName, dataflowProfileName, dataflowName);
IotOperationsDataflowResource iotOperationsDataflow = client.GetIotOperationsDataflowResource(iotOperationsDataflowResourceId);

// invoke the operation
IotOperationsDataflowResource result = await iotOperationsDataflow.GetAsync();

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
IotOperationsDataflowData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
