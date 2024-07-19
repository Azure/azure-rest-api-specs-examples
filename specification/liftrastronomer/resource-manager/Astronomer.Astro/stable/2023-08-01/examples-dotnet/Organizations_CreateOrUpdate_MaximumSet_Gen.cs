using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Astro.Models;
using Azure.ResourceManager.Models;
using Azure.ResourceManager.Resources;
using Azure.ResourceManager.Astro;

// Generated from example definition: specification/liftrastronomer/resource-manager/Astronomer.Astro/stable/2023-08-01/examples/Organizations_CreateOrUpdate_MaximumSet_Gen.json
// this example is just showing the usage of "Organizations_CreateOrUpdate" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ResourceGroupResource created on azure
// for more information of creating ResourceGroupResource, please refer to the document of ResourceGroupResource
string subscriptionId = "43454B17-172A-40FE-80FA-549EA23D12B3";
string resourceGroupName = "rgastronomer";
ResourceIdentifier resourceGroupResourceId = ResourceGroupResource.CreateResourceIdentifier(subscriptionId, resourceGroupName);
ResourceGroupResource resourceGroupResource = client.GetResourceGroupResource(resourceGroupResourceId);

// get the collection of this AstroOrganizationResource
AstroOrganizationCollection collection = resourceGroupResource.GetAstroOrganizations();

// invoke the operation
string organizationName = "U.1-:7";
AstroOrganizationData data = new AstroOrganizationData(new AzureLocation("mhqthlsatwvqkl"))
{
    Properties = new AstroOrganizationProperties(new AstroMarketplaceDetails("ntthclydlpqmasr", new AstroOfferDetails("gfsqxygpnerxmvols", "krzkefmpxztqyusidzgpchfaswuyce", "kndxzygsanuiqzwbfbbvoipv")
    {
        PlanName = "pwqjwlq",
        TermUnit = "xyygyzcazkuelz",
        TermId = "pwds",
    })
    {
        SubscriptionStatus = MarketplaceSubscriptionStatus.PendingFulfillmentStart,
    }, new AstroUserDetails("nfh", "lazfbstcccykibvcrxpmglqam", ".K_@e7N-g1.xjqnbPs")
    {
        Upn = "xtutvycpxjrtoftx",
        PhoneNumber = "inxkscllh",
    })
    {
        PartnerOrganizationProperties = new AstroPartnerOrganizationProperties("3-")
        {
            OrganizationId = "lskgzdmziusgrsucv",
            WorkspaceId = "vcrupxwpaba",
            WorkspaceName = "9.:06",
            SingleSignOnProperties = new AstroSingleSignOnProperties()
            {
                SingleSignOnState = AstroSingleSignOnState.Initial,
                EnterpriseAppId = "mklfypyujwumgwdzae",
                SingleSignOnUri = new Uri("ymmtzkyghvinvhgnqlzwrr"),
                AadDomains =
                {
                "kfbleh"
                },
            },
        },
    },
    Identity = new ManagedServiceIdentity("None")
    {
        UserAssignedIdentities =
        {
        },
    },
    Tags =
    {
    },
};
ArmOperation<AstroOrganizationResource> lro = await collection.CreateOrUpdateAsync(WaitUntil.Completed, organizationName, data);
AstroOrganizationResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
AstroOrganizationData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
