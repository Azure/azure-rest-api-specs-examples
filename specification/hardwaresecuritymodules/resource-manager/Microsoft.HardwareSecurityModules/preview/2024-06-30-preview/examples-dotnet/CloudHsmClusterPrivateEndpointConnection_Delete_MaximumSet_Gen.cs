using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.HardwareSecurityModules.Models;
using Azure.ResourceManager.HardwareSecurityModules;

// Generated from example definition: specification/hardwaresecuritymodules/resource-manager/Microsoft.HardwareSecurityModules/preview/2024-06-30-preview/examples/CloudHsmClusterPrivateEndpointConnection_Delete_MaximumSet_Gen.json
// this example is just showing the usage of "CloudHsmClusterPrivateEndpointConnections_Delete" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this CloudHsmClusterPrivateEndpointConnectionResource created on azure
// for more information of creating CloudHsmClusterPrivateEndpointConnectionResource, please refer to the document of CloudHsmClusterPrivateEndpointConnectionResource
string subscriptionId = "00000000-0000-0000-0000-000000000000";
string resourceGroupName = "rgcloudhsm";
string cloudHsmClusterName = "chsm1";
string peConnectionName = "sample-pec";
ResourceIdentifier cloudHsmClusterPrivateEndpointConnectionResourceId = CloudHsmClusterPrivateEndpointConnectionResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, cloudHsmClusterName, peConnectionName);
CloudHsmClusterPrivateEndpointConnectionResource cloudHsmClusterPrivateEndpointConnection = client.GetCloudHsmClusterPrivateEndpointConnectionResource(cloudHsmClusterPrivateEndpointConnectionResourceId);

// invoke the operation
await cloudHsmClusterPrivateEndpointConnection.DeleteAsync(WaitUntil.Completed);

Console.WriteLine($"Succeeded");
