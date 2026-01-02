using Azure;
using Azure.ResourceManager;
using System;
using System.Collections.Generic;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.IotOperations.Models;
using Azure.ResourceManager.IotOperations;

// Generated from example definition: 2025-07-01-preview/BrokerAuthentication_CreateOrUpdate_Complex.json
// this example is just showing the usage of "BrokerAuthenticationResource_CreateOrUpdate" operation, for the dependent resources, they will have to be created separately.

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

// get the collection of this IotOperationsBrokerAuthenticationResource
IotOperationsBrokerAuthenticationCollection collection = iotOperationsBroker.GetIotOperationsBrokerAuthentications();

// invoke the operation
string authenticationName = "resource-name123";
IotOperationsBrokerAuthenticationData data = new IotOperationsBrokerAuthenticationData
{
    Properties = new IotOperationsBrokerAuthenticationProperties(new BrokerAuthenticatorMethods[]
{
new BrokerAuthenticatorMethods(BrokerAuthenticationMethod.ServiceAccountToken)
{
ServiceAccountTokenAudiences = {"aio-internal"},
},
new BrokerAuthenticatorMethods(BrokerAuthenticationMethod.X509)
{
X509Settings = new BrokerAuthenticatorMethodX509
{
AuthorizationAttributes =
{
["root"] = new BrokerAuthenticatorMethodX509Attributes(new Dictionary<string, string>
{
["organization"] = "contoso"
}, "CN = Contoso Root CA Cert, OU = Engineering, C = US"),
["intermediate"] = new BrokerAuthenticatorMethodX509Attributes(new Dictionary<string, string>
{
["city"] = "seattle",
["foo"] = "bar"
}, "CN = Contoso Intermediate CA"),
["smart-fan"] = new BrokerAuthenticatorMethodX509Attributes(new Dictionary<string, string>
{
["building"] = "17"
}, "CN = smart-fan")
},
TrustedClientCaCert = "my-ca",
},
}
}),
    ExtendedLocation = new IotOperationsExtendedLocation("/subscriptions/F8C729F9-DF9C-4743-848F-96EE433D8E53/resourceGroups/rgiotoperations/providers/Microsoft.ExtendedLocation/customLocations/resource-123", IotOperationsExtendedLocationType.CustomLocation),
};
ArmOperation<IotOperationsBrokerAuthenticationResource> lro = await collection.CreateOrUpdateAsync(WaitUntil.Completed, authenticationName, data);
IotOperationsBrokerAuthenticationResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
IotOperationsBrokerAuthenticationData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
