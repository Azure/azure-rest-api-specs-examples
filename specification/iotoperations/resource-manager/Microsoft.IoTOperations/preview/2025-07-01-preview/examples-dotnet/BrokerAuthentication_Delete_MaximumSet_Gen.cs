using Azure;
using Azure.ResourceManager;
using System;
using System.Collections.Generic;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.IotOperations.Models;
using Azure.ResourceManager.IotOperations;

// Generated from example definition: 2025-07-01-preview/BrokerAuthentication_Delete_MaximumSet_Gen.json
// this example is just showing the usage of "BrokerAuthenticationResource_Delete" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this IotOperationsBrokerAuthenticationResource created on azure
// for more information of creating IotOperationsBrokerAuthenticationResource, please refer to the document of IotOperationsBrokerAuthenticationResource
string subscriptionId = "F8C729F9-DF9C-4743-848F-96EE433D8E53";
string resourceGroupName = "rgiotoperations";
string instanceName = "resource-name123";
string brokerName = "resource-name123";
string authenticationName = "resource-name123";
ResourceIdentifier iotOperationsBrokerAuthenticationResourceId = IotOperationsBrokerAuthenticationResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, instanceName, brokerName, authenticationName);
IotOperationsBrokerAuthenticationResource iotOperationsBrokerAuthentication = client.GetIotOperationsBrokerAuthenticationResource(iotOperationsBrokerAuthenticationResourceId);

// invoke the operation
await iotOperationsBrokerAuthentication.DeleteAsync(WaitUntil.Completed);

Console.WriteLine("Succeeded");
