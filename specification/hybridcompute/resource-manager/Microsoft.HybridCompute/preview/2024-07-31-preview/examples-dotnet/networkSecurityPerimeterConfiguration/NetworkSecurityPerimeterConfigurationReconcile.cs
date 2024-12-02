using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.HybridCompute.Models;
using Azure.ResourceManager.HybridCompute;

// Generated from example definition: specification/hybridcompute/resource-manager/Microsoft.HybridCompute/preview/2024-07-31-preview/examples/networkSecurityPerimeterConfiguration/NetworkSecurityPerimeterConfigurationReconcile.json
// this example is just showing the usage of "NetworkSecurityPerimeterConfigurations_ReconcileForPrivateLinkScope" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this NetworkSecurityPerimeterConfigurationResource created on azure
// for more information of creating NetworkSecurityPerimeterConfigurationResource, please refer to the document of NetworkSecurityPerimeterConfigurationResource
string subscriptionId = "00000000-1111-2222-3333-444444444444";
string resourceGroupName = "my-resource-group";
string scopeName = "my-privatelinkscope";
string perimeterName = "aaaaaaaa-bbbb-cccc-dddd-eeeeeeeeeeee.myAssociation";
ResourceIdentifier networkSecurityPerimeterConfigurationResourceId = NetworkSecurityPerimeterConfigurationResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, scopeName, perimeterName);
NetworkSecurityPerimeterConfigurationResource networkSecurityPerimeterConfiguration = client.GetNetworkSecurityPerimeterConfigurationResource(networkSecurityPerimeterConfigurationResourceId);

// invoke the operation
ArmOperation<NetworkSecurityPerimeterConfigurationReconcileResult> lro = await networkSecurityPerimeterConfiguration.ReconcileForPrivateLinkScopeAsync(WaitUntil.Completed);
NetworkSecurityPerimeterConfigurationReconcileResult result = lro.Value;

Console.WriteLine($"Succeeded: {result}");
