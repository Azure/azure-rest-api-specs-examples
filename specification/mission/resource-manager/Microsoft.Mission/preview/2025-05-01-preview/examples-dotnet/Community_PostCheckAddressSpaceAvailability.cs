using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.VirtualEnclaves.Models;
using Azure.ResourceManager.VirtualEnclaves;

// Generated from example definition: 2025-05-01-preview/Community_PostCheckAddressSpaceAvailability.json
// this example is just showing the usage of "Community_CheckAddressSpaceAvailability" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this VirtualEnclaveCommunityResource created on azure
// for more information of creating VirtualEnclaveCommunityResource, please refer to the document of VirtualEnclaveCommunityResource
string subscriptionId = "CA1CB369-DD26-4DB2-9D43-9AFEF0F22093";
string resourceGroupName = "rgopenapi";
string communityName = "TestMyCommunity";
ResourceIdentifier virtualEnclaveCommunityResourceId = VirtualEnclaveCommunityResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, communityName);
VirtualEnclaveCommunityResource virtualEnclaveCommunity = client.GetVirtualEnclaveCommunityResource(virtualEnclaveCommunityResourceId);

// invoke the operation
CheckAddressSpaceAvailabilityContent content = new CheckAddressSpaceAvailabilityContent(new ResourceIdentifier("/subscriptions/c64f6eca-bdc5-4bc2-88d6-f8f1dc23f86c/resourceGroups/TestMyRg/providers/Microsoft.Mission/communities/TestMyCommunity"), new EnclaveVirtualNetwork
{
    NetworkSize = "small",
    CustomCidrRange = "10.0.0.0/24",
    SubnetConfigurations = { new VirtualEnclaveSubnetConfiguration("test", 26) },
    AllowSubnetCommunication = true,
});
CheckAddressSpaceAvailabilityResult result = await virtualEnclaveCommunity.CheckAddressSpaceAvailabilityAsync(content);

Console.WriteLine($"Succeeded: {result}");
