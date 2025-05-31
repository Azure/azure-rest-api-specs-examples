using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Models;
using Azure.ResourceManager.MongoDBAtlas.Models;
using Azure.ResourceManager.Resources;
using Azure.ResourceManager.MongoDBAtlas;

// Generated from example definition: 2024-11-18-preview/Organizations_CreateOrUpdate_MaximumSet_Gen.json
// this example is just showing the usage of "OrganizationResource_CreateOrUpdate" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ResourceGroupResource created on azure
// for more information of creating ResourceGroupResource, please refer to the document of ResourceGroupResource
string subscriptionId = "AD4FE133-6EF1-4ED8-82DB-5C1CBA58597E";
string resourceGroupName = "rgopenapi";
ResourceIdentifier resourceGroupResourceId = ResourceGroupResource.CreateResourceIdentifier(subscriptionId, resourceGroupName);
ResourceGroupResource resourceGroupResource = client.GetResourceGroupResource(resourceGroupResourceId);

// get the collection of this MongoDBAtlasOrganizationResource
MongoDBAtlasOrganizationCollection collection = resourceGroupResource.GetMongoDBAtlasOrganizations();

// invoke the operation
string organizationName = "U.1-:7";
MongoDBAtlasOrganizationData data = new MongoDBAtlasOrganizationData(new AzureLocation("wobqn"))
{
    Properties = new MongoDBAtlasOrganizationProperties(new MongoDBAtlasMarketplaceDetails("o", new MongoDBAtlasOfferDetails("rxglearenxsgpwzlsxmiicynks", "ohnquleylybvjrtnpjupvwlk", "obhxnhvrtbcnoovgofbs")
    {
        PlanName = "lkwdzpfhvjezjusrqzyftcikxdt",
        TermUnit = "omkxrnburbnruglwqgjlahvjmbfcse",
        TermId = "bqmmltwmtpdcdeszbka",
    }), new MongoDBAtlasUserDetails("aslybvdwwddqxwazxvxhjrs", "cnuitqoqpcyvmuqowgnxpwxjcveyr", ".K_@e7N-g1.xjqnbPs")
    {
        Upn = "howdzmfy",
        PhoneNumber = "ilypntsrbmbbbexbasuu",
        CompanyName = "oxdcwwl",
    })
    {
        PartnerProperties = new MongoDBAtlasPartnerProperties("U.1-:7")
        {
            OrganizationId = "lyombjlhvwxithkiy",
            RedirectUri = "cbxwtehraetlluocdihfgchvjzockn",
        },
    },
    Identity = new ManagedServiceIdentity("None")
    {
        UserAssignedIdentities = { },
    },
    Tags = { },
};
ArmOperation<MongoDBAtlasOrganizationResource> lro = await collection.CreateOrUpdateAsync(WaitUntil.Completed, organizationName, data);
MongoDBAtlasOrganizationResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
MongoDBAtlasOrganizationData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
