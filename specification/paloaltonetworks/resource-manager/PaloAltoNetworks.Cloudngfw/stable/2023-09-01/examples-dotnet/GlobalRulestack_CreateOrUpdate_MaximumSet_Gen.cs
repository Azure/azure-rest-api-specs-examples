using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Models;
using Azure.ResourceManager.PaloAltoNetworks.Ngfw;
using Azure.ResourceManager.PaloAltoNetworks.Ngfw.Models;

// Generated from example definition: specification/paloaltonetworks/resource-manager/PaloAltoNetworks.Cloudngfw/stable/2023-09-01/examples/GlobalRulestack_CreateOrUpdate_MaximumSet_Gen.json
// this example is just showing the usage of "GlobalRulestack_CreateOrUpdate" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this TenantResource created on azure
// for more information of creating TenantResource, please refer to the document of TenantResource
var tenantResource = client.GetTenants().GetAllAsync().GetAsyncEnumerator().Current;

// get the collection of this GlobalRulestackResource
GlobalRulestackCollection collection = tenantResource.GetGlobalRulestacks();

// invoke the operation
string globalRulestackName = "praval";
GlobalRulestackData data = new GlobalRulestackData(new AzureLocation("eastus"))
{
    Identity = new ManagedServiceIdentity("None")
    {
        UserAssignedIdentities =
        {
        [new ResourceIdentifier("key16")] = new UserAssignedIdentity(),
        },
    },
    PanETag = new ETag("2bf4a339-294d-4c25-b0b2-ef649e9f5c12"),
    PanLocation = new AzureLocation("eastus"),
    Scope = RulestackScopeType.Global,
    AssociatedSubscriptions =
    {
    "2bf4a339-294d-4c25-b0b2-ef649e9f5c27"
    },
    Description = "global rulestacks",
    DefaultMode = RuleCreationDefaultMode.IPS,
    MinAppIdVersion = "8.5.3",
    SecurityServices = new RulestackSecurityServices()
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
};
ArmOperation<GlobalRulestackResource> lro = await collection.CreateOrUpdateAsync(WaitUntil.Completed, globalRulestackName, data);
GlobalRulestackResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
GlobalRulestackData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
