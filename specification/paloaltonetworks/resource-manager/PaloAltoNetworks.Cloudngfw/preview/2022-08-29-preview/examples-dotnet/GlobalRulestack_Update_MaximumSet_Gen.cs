using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Models;
using Azure.ResourceManager.PaloAltoNetworks.Ngfw;
using Azure.ResourceManager.PaloAltoNetworks.Ngfw.Models;

// Generated from example definition: specification/paloaltonetworks/resource-manager/PaloAltoNetworks.Cloudngfw/preview/2022-08-29-preview/examples/GlobalRulestack_Update_MaximumSet_Gen.json
// this example is just showing the usage of "GlobalRulestack_Update" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this GlobalRulestackResource created on azure
// for more information of creating GlobalRulestackResource, please refer to the document of GlobalRulestackResource
string globalRulestackName = "praval";
ResourceIdentifier globalRulestackResourceId = GlobalRulestackResource.CreateResourceIdentifier(globalRulestackName);
GlobalRulestackResource globalRulestackResource = client.GetGlobalRulestackResource(globalRulestackResourceId);

// invoke the operation
GlobalRulestackResourcePatch patch = new GlobalRulestackResourcePatch()
{
    Location = new AzureLocation("eastus"),
    Identity = new AzureResourceManagerManagedIdentityProperties(ManagedIdentityType.None)
    {
        UserAssignedIdentities =
        {
        ["key16"] = new UserAssignedIdentity(),
        },
    },
    Properties = new GlobalRulestackResourceUpdateProperties()
    {
        PanETag = "2bf4a339-294d-4c25-b0b2-ef649e9f5c12",
        PanLocation = "eastus",
        Scope = ScopeType.Global,
        AssociatedSubscriptions =
        {
        "2bf4a339-294d-4c25-b0b2-ef649e9f5c27"
        },
        Description = "global rulestacks",
        DefaultMode = DefaultMode.IPS,
        MinAppIdVersion = "8.5.3",
        SecurityServices = new SecurityServices()
        {
            VulnerabilityProfile = "default",
            AntiSpywareProfile = "default",
            AntiVirusProfile = "default",
            UrlFilteringProfile = "default",
            FileBlockingProfile = "default",
            DnsSubscription = "default",
            OutboundUnTrustCertificate = "default",
            OutboundTrustCertificate = "default",
        },
    },
};
GlobalRulestackResource result = await globalRulestackResource.UpdateAsync(patch);

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
GlobalRulestackResourceData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
