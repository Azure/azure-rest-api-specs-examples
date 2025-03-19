using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Models;
using Azure.ResourceManager.PaloAltoNetworks.Ngfw.Models;
using Azure.ResourceManager.Resources;
using Azure.ResourceManager.PaloAltoNetworks.Ngfw;

// Generated from example definition: specification/paloaltonetworks/resource-manager/PaloAltoNetworks.Cloudngfw/stable/2023-09-01/examples/LocalRulestacks_CreateOrUpdate_MaximumSet_Gen.json
// this example is just showing the usage of "LocalRulestacks_CreateOrUpdate" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ResourceGroupResource created on azure
// for more information of creating ResourceGroupResource, please refer to the document of ResourceGroupResource
string subscriptionId = "2bf4a339-294d-4c25-b0b2-ef649e9f5c27";
string resourceGroupName = "rgopenapi";
ResourceIdentifier resourceGroupResourceId = ResourceGroupResource.CreateResourceIdentifier(subscriptionId, resourceGroupName);
ResourceGroupResource resourceGroupResource = client.GetResourceGroupResource(resourceGroupResourceId);

// get the collection of this LocalRulestackResource
LocalRulestackCollection collection = resourceGroupResource.GetLocalRulestacks();

// invoke the operation
string localRulestackName = "lrs1";
LocalRulestackData data = new LocalRulestackData(new AzureLocation("eastus"))
{
    Identity = new ManagedServiceIdentity("None")
    {
        UserAssignedIdentities =
        {
        [new ResourceIdentifier("key16")] = new UserAssignedIdentity()
        },
    },
    PanETag = new ETag("2bf4a339-294d-4c25-b0b2-ef649e9f5c12"),
    PanLocation = new AzureLocation("eastus"),
    Scope = RulestackScopeType.Local,
    AssociatedSubscriptions = { "2bf4a339-294d-4c25-b0b2-ef649e9f5c27" },
    Description = "local rulestacks",
    DefaultMode = RuleCreationDefaultMode.IPS,
    MinAppIdVersion = "8.5.3",
    SecurityServices = new RulestackSecurityServices
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
    Tags =
    {
    ["tagName"] = "value"
    },
};
ArmOperation<LocalRulestackResource> lro = await collection.CreateOrUpdateAsync(WaitUntil.Completed, localRulestackName, data);
LocalRulestackResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
LocalRulestackData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
