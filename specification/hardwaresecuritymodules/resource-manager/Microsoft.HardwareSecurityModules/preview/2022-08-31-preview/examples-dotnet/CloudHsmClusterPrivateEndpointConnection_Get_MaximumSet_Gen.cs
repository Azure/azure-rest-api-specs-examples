using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.HardwareSecurityModules;
using Azure.ResourceManager.HardwareSecurityModules.Models;

// Generated from example definition: specification/hardwaresecuritymodules/resource-manager/Microsoft.HardwareSecurityModules/preview/2022-08-31-preview/examples/CloudHsmClusterPrivateEndpointConnection_Get_MaximumSet_Gen.json
// this example is just showing the usage of "CloudHsmClusterPrivateEndpointConnections_Get" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this HardwareSecurityModulesPrivateEndpointConnectionResource created on azure
// for more information of creating HardwareSecurityModulesPrivateEndpointConnectionResource, please refer to the document of HardwareSecurityModulesPrivateEndpointConnectionResource
string subscriptionId = "00000000-0000-0000-0000-000000000000";
string resourceGroupName = "rgcloudhsm";
string cloudHsmClusterName = "chsm1";
string peConnectionName = "sample-pec";
ResourceIdentifier hardwareSecurityModulesPrivateEndpointConnectionResourceId = HardwareSecurityModulesPrivateEndpointConnectionResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, cloudHsmClusterName, peConnectionName);
HardwareSecurityModulesPrivateEndpointConnectionResource hardwareSecurityModulesPrivateEndpointConnection = client.GetHardwareSecurityModulesPrivateEndpointConnectionResource(hardwareSecurityModulesPrivateEndpointConnectionResourceId);

// invoke the operation
HardwareSecurityModulesPrivateEndpointConnectionResource result = await hardwareSecurityModulesPrivateEndpointConnection.GetAsync();

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
HardwareSecurityModulesPrivateEndpointConnectionData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
