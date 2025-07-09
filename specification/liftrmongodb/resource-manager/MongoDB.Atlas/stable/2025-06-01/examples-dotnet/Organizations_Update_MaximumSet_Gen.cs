using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Models;
using Azure.ResourceManager.MongoDBAtlas.Models;
using Azure.ResourceManager.MongoDBAtlas;

// Generated from example definition: 2025-06-01/Organizations_Update_MaximumSet_Gen.json
// this example is just showing the usage of "OrganizationResource_Update" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this MongoDBAtlasOrganizationResource created on azure
// for more information of creating MongoDBAtlasOrganizationResource, please refer to the document of MongoDBAtlasOrganizationResource
string subscriptionId = "422A4D59-A5BC-4DBB-8831-EC666633F64F";
string resourceGroupName = "rgopenapi";
string organizationName = "U.1-:7";
ResourceIdentifier mongoDBAtlasOrganizationResourceId = MongoDBAtlasOrganizationResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, organizationName);
MongoDBAtlasOrganizationResource mongoDBAtlasOrganization = client.GetMongoDBAtlasOrganizationResource(mongoDBAtlasOrganizationResourceId);

// invoke the operation
MongoDBAtlasOrganizationPatch patch = new MongoDBAtlasOrganizationPatch
{
    Identity = new ManagedServiceIdentity("None")
    {
        UserAssignedIdentities = { },
    },
    Tags = { },
    Properties = new MongoDBAtlasOrganizationUpdateProperties
    {
        User = new MongoDBAtlasUserDetails("btyhwmlbzzihjfimviefebg", "xx", ".K_@e7N-g1.xjqnbPs")
        {
            Upn = "mxtbogd",
            PhoneNumber = "isvc",
            CompanyName = "oztteysco",
        },
        PartnerProperties = new MongoDBAtlasPartnerProperties("U.1-:7")
        {
            OrganizationId = "vugtqrobendjkinziswxlqueouo",
            RedirectUri = "cbxwtehraetlluocdihfgchvjzockn",
        },
    },
};
ArmOperation<MongoDBAtlasOrganizationResource> lro = await mongoDBAtlasOrganization.UpdateAsync(WaitUntil.Completed, patch);
MongoDBAtlasOrganizationResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
MongoDBAtlasOrganizationData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
