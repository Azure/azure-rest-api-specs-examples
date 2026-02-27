using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Batch;

// Generated from example definition: specification/batch/resource-manager/Microsoft.Batch/Batch/stable/2025-06-01/examples/NspConfigurationReconcile.json
// this example is just showing the usage of "NetworkSecurityPerimeter_ReconcileConfiguration" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this NetworkSecurityPerimeterConfigurationResource created on azure
// for more information of creating NetworkSecurityPerimeterConfigurationResource, please refer to the document of NetworkSecurityPerimeterConfigurationResource
string subscriptionId = "12345678-1234-1234-1234-123456789012";
string resourceGroupName = "default-azurebatch-japaneast";
string accountName = "sampleacct";
string networkSecurityPerimeterConfigurationName = "00000000-0000-0000-0000-000000000000.sampleassociation";
ResourceIdentifier networkSecurityPerimeterConfigurationResourceId = NetworkSecurityPerimeterConfigurationResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, accountName, networkSecurityPerimeterConfigurationName);
NetworkSecurityPerimeterConfigurationResource networkSecurityPerimeterConfiguration = client.GetNetworkSecurityPerimeterConfigurationResource(networkSecurityPerimeterConfigurationResourceId);

// invoke the operation
await networkSecurityPerimeterConfiguration.ReconcileConfigurationAsync(WaitUntil.Completed);

Console.WriteLine("Succeeded");
