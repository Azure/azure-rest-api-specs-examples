using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.NetworkCloud.Models;
using Azure.ResourceManager.NetworkCloud;

// Generated from example definition: specification/networkcloud/resource-manager/Microsoft.NetworkCloud/stable/2025-02-01/examples/BareMetalMachines_Cordon.json
// this example is just showing the usage of "BareMetalMachines_Cordon" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this NetworkCloudBareMetalMachineResource created on azure
// for more information of creating NetworkCloudBareMetalMachineResource, please refer to the document of NetworkCloudBareMetalMachineResource
string subscriptionId = "123e4567-e89b-12d3-a456-426655440000";
string resourceGroupName = "resourceGroupName";
string bareMetalMachineName = "bareMetalMachineName";
ResourceIdentifier networkCloudBareMetalMachineResourceId = NetworkCloudBareMetalMachineResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, bareMetalMachineName);
NetworkCloudBareMetalMachineResource networkCloudBareMetalMachine = client.GetNetworkCloudBareMetalMachineResource(networkCloudBareMetalMachineResourceId);

// invoke the operation
BareMetalMachineCordonContent content = new BareMetalMachineCordonContent
{
    Evacuate = BareMetalMachineEvacuate.True,
};
ArmOperation<NetworkCloudOperationStatusResult> lro = await networkCloudBareMetalMachine.CordonAsync(WaitUntil.Completed, content: content);
NetworkCloudOperationStatusResult result = lro.Value;

Console.WriteLine($"Succeeded: {result}");
