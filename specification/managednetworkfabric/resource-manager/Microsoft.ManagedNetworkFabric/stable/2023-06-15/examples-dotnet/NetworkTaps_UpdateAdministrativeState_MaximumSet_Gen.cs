using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.ManagedNetworkFabric;
using Azure.ResourceManager.ManagedNetworkFabric.Models;
using Azure.ResourceManager.Resources;

// Generated from example definition: specification/managednetworkfabric/resource-manager/Microsoft.ManagedNetworkFabric/stable/2023-06-15/examples/NetworkTaps_UpdateAdministrativeState_MaximumSet_Gen.json
// this example is just showing the usage of "NetworkTaps_UpdateAdministrativeState" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this NetworkTapResource created on azure
// for more information of creating NetworkTapResource, please refer to the document of NetworkTapResource
string subscriptionId = "1234ABCD-0A1B-1234-5678-123456ABCDEF";
string resourceGroupName = "example-rg";
string networkTapName = "example-networkTap";
ResourceIdentifier networkTapResourceId = NetworkTapResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, networkTapName);
NetworkTapResource networkTap = client.GetNetworkTapResource(networkTapResourceId);

// invoke the operation
UpdateAdministrativeStateContent content = new UpdateAdministrativeStateContent()
{
    State = AdministrativeEnableState.Enable,
    ResourceIds =
    {
    new ResourceIdentifier("")
    },
};
ArmOperation<DeviceUpdateCommonPostActionResult> lro = await networkTap.UpdateAdministrativeStateAsync(WaitUntil.Completed, content);
DeviceUpdateCommonPostActionResult result = lro.Value;

Console.WriteLine($"Succeeded: {result}");
