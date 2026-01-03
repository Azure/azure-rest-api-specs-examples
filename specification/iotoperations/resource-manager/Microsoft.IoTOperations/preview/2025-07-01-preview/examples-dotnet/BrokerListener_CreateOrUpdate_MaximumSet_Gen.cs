using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.IotOperations.Models;
using Azure.ResourceManager.IotOperations;

// Generated from example definition: 2025-07-01-preview/BrokerListener_CreateOrUpdate_MaximumSet_Gen.json
// this example is just showing the usage of "BrokerListenerResource_CreateOrUpdate" operation, for the dependent resources, they will have to be created separately.

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

// get the collection of this IotOperationsBrokerListenerResource
IotOperationsBrokerListenerCollection collection = iotOperationsBroker.GetIotOperationsBrokerListeners();

// invoke the operation
string listenerName = "resource-name123";
IotOperationsBrokerListenerData data = new IotOperationsBrokerListenerData
{
    Properties = new IotOperationsBrokerListenerProperties(new BrokerListenerPort[]
{
new BrokerListenerPort(1268)
{
AuthenticationRef = "tjvdroaqqy",
AuthorizationRef = "inxhvxnwswyrvt",
NodePort = 7281,
Protocol = BrokerProtocolType.Mqtt,
Tls = new ListenerPortTlsCertMethod(TlsCertMethodMode.Automatic)
{
CertManagerCertificateSpec = new CertManagerCertificateSpec(new CertManagerIssuerRef("jtmuladdkpasfpoyvewekmiy", CertManagerIssuerKind.Issuer, "ocwoqpgucvjrsuudtjhb"))
{
Duration = "qmpeffoksron",
SecretName = "oagi",
RenewBefore = "hutno",
PrivateKey = new CertManagerPrivateKey(PrivateKeyAlgorithm.Ec256, PrivateKeyRotationPolicy.Always),
San = new SanForCert(new string[]{"xhvmhrrhgfsapocjeebqtnzarlj"}, new string[]{"zbgugfzcgsmegevzktsnibyuyp"}),
},
ManualSecretRef = "secret-name",
},
}
})
    {
        ServiceName = "tpfiszlapdpxktx",
        ServiceType = BlockerListenerServiceType.ClusterIP,
    },
    ExtendedLocation = new IotOperationsExtendedLocation("/subscriptions/F8C729F9-DF9C-4743-848F-96EE433D8E53/resourceGroups/rgiotoperations/providers/Microsoft.ExtendedLocation/customLocations/resource-123", IotOperationsExtendedLocationType.CustomLocation),
};
ArmOperation<IotOperationsBrokerListenerResource> lro = await collection.CreateOrUpdateAsync(WaitUntil.Completed, listenerName, data);
IotOperationsBrokerListenerResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
IotOperationsBrokerListenerData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
