using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.ManagedNetworkFabric.Models;
using Azure.ResourceManager.ManagedNetworkFabric;

// Generated from example definition: specification/managednetworkfabric/resource-manager/Microsoft.ManagedNetworkFabric/stable/2023-06-15/examples/L2IsolationDomains_updateAdministrativeState_MaximumSet_Gen.json
// this example is just showing the usage of "L2IsolationDomains_UpdateAdministrativeState" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this NetworkFabricL2IsolationDomainResource created on azure
// for more information of creating NetworkFabricL2IsolationDomainResource, please refer to the document of NetworkFabricL2IsolationDomainResource
string subscriptionId = "1234ABCD-0A1B-1234-5678-123456ABCDEF";
string resourceGroupName = "example-rg";
string l2IsolationDomainName = "example-l2Domain";
ResourceIdentifier networkFabricL2IsolationDomainResourceId = NetworkFabricL2IsolationDomainResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, l2IsolationDomainName);
NetworkFabricL2IsolationDomainResource networkFabricL2IsolationDomain = client.GetNetworkFabricL2IsolationDomainResource(networkFabricL2IsolationDomainResourceId);

// invoke the operation
UpdateAdministrativeStateContent content = new UpdateAdministrativeStateContent
{
    State = AdministrativeEnableState.Enable,
    ResourceIds = { new ResourceIdentifier("") },
};
ArmOperation<DeviceUpdateCommonPostActionResult> lro = await networkFabricL2IsolationDomain.UpdateAdministrativeStateAsync(WaitUntil.Completed, content);
DeviceUpdateCommonPostActionResult result = lro.Value;

Console.WriteLine($"Succeeded: {result}");
