using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Astro.Models;
using Azure.ResourceManager.Models;
using Azure.ResourceManager.Astro;

// Generated from example definition: specification/liftrastronomer/resource-manager/Astronomer.Astro/stable/2023-08-01/examples/Organizations_Update_MaximumSet_Gen.json
// this example is just showing the usage of "Organizations_Update" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this AstroOrganizationResource created on azure
// for more information of creating AstroOrganizationResource, please refer to the document of AstroOrganizationResource
string subscriptionId = "43454B17-172A-40FE-80FA-549EA23D12B3";
string resourceGroupName = "rgastronomer";
string organizationName = "6.";
ResourceIdentifier astroOrganizationResourceId = AstroOrganizationResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, organizationName);
AstroOrganizationResource astroOrganization = client.GetAstroOrganizationResource(astroOrganizationResourceId);

// invoke the operation
AstroOrganizationPatch patch = new AstroOrganizationPatch
{
    Identity = new ManagedServiceIdentity("None")
    {
        UserAssignedIdentities = { },
    },
    Tags =
    {
    ["key1474"] = "bqqyipxnbbxryhznyaosmtpo"
    },
    Properties = new AstroOrganizationUpdateProperties
    {
        User = new AstroUserUpdateDetails
        {
            FirstName = "qeuofehzypzljgcuysugefbgxde",
            LastName = "g",
            EmailAddress = ".K_@e7N-g1.xjqnbPs",
            Upn = "uwtprzdfpsqmktx",
            PhoneNumber = "aqpyxznvqpgkzohevynofrjdfgoo",
        },
        PartnerOrganizationProperties = new AstroPartnerOrganizationUpdateProperties
        {
            OrganizationId = "lrtmbkvyvvoszhjevohkmyjhfyty",
            WorkspaceId = "xsepuskdhejaadusyxq",
            OrganizationName = "U2P_",
            WorkspaceName = "L.-y_--:",
            SingleSignOnProperties = new AstroSingleSignOnProperties
            {
                SingleSignOnState = AstroSingleSignOnState.Initial,
                EnterpriseAppId = "mklfypyujwumgwdzae",
                SingleSignOnUri = new Uri("ymmtzkyghvinvhgnqlzwrr"),
                AadDomains = { "kfbleh" },
            },
        },
    },
};
ArmOperation<AstroOrganizationResource> lro = await astroOrganization.UpdateAsync(WaitUntil.Completed, patch);
AstroOrganizationResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
AstroOrganizationData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
