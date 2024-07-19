using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.InformaticaDataManagement.Models;
using Azure.ResourceManager.Resources;
using Azure.ResourceManager.InformaticaDataManagement;

// Generated from example definition: specification/informatica/resource-manager/Informatica.DataManagement/stable/2024-05-08/examples/Organizations_Update_MaximumSet_Gen.json
// this example is just showing the usage of "Organizations_Update" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this InformaticaOrganizationResource created on azure
// for more information of creating InformaticaOrganizationResource, please refer to the document of InformaticaOrganizationResource
string subscriptionId = "3599DA28-E346-4D9F-811E-189C0445F0FE";
string resourceGroupName = "rgopenapi";
string organizationName = "_-";
ResourceIdentifier informaticaOrganizationResourceId = InformaticaOrganizationResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, organizationName);
InformaticaOrganizationResource informaticaOrganization = client.GetInformaticaOrganizationResource(informaticaOrganizationResourceId);

// invoke the operation
InformaticaOrganizationPatch patch = new InformaticaOrganizationPatch()
{
    Tags =
    {
    ["key1918"] = "fbjvtuvzsghpl",
    },
    Properties = new InformaticaOrganizationPropertiesUpdate()
    {
        MarketplaceDetails = new InformaticaMarketplaceDetailsUpdate()
        {
            MarketplaceSubscriptionId = "szhyxzgjtssjmlguivepc",
            OfferDetails = new InformaticaOfferDetailsUpdate()
            {
                PublisherId = "ktzfghsyjqbsswhltoaemgotmnorhdogvkaxplutbjjqzuepxizliynyakersobagvpwvpzwjtjjxigsqgcyqaahaxdijghnexliofhfjlqzjmmbvrhcvjxdodnexxizbgfhjopbwzjojxsluasnwwsgcajefglbcvzpaeblanhmurcculndtfwnfjyxol",
                OfferId = "idaxbflabvjsippplyenvrpgeydsjxcmyubgukffkcdvlvrtwpdhnvdblxjsldiuswrchsibk",
                PlanId = "giihvvnwdwzkfqrhkpqzbgfotzyixnsvmxzauseebillhslauglzfxzvzvts",
                PlanName = "tfqjenotaewzdeerliteqxdawuqxhwdzbtiiimsaedrlsnbdoonnloakjtvnwhhrcyxxsgoachguthqvlahpjyofpoqpfacfmiaauawazkmxkjgvktbptojknzojtjrfzvbbjjkvstabqyaczxinijhoxrjukftsagpwgsvpmczopztmplipyufhuaumfx",
                TermUnit = "nykqoplazujcwmfldntifjqrnx",
                TermId = "eolmwogtgpdncqoigqcdomupwummaicwvdxgbskpdsmjizdfbdgbxbuekcpwmenqzbhqxpdnjtup",
            },
        },
        UserDetails = new InformaticaUserDetailsUpdate()
        {
            FirstName = "qguqrmanyupoi",
            LastName = "ugzg",
            EmailAddress = "7_-46@13D--3.m-4x-.11.c-9-.DHLYFc",
            Upn = "viwjrkn",
            PhoneNumber = "uxa",
        },
        CompanyDetails = new InformaticaCompanyDetailsUpdate()
        {
            CompanyName = "xkrvbozrjcvappqeeyt",
            OfficeAddress = "sfcx",
            Country = "rvlzppgvopcw",
            Domain = "dponvwnrdrnzahcurqssesukbsokdd",
            Business = "mwqblnruflwpolgbxpqbqneve",
            NumberOfEmployees = 22,
        },
        ExistingResourceId = new ResourceIdentifier("/subscriptions/subid/resourceGroups/rg1/providers/Informatica.DataManagement/organizations/org1/serverlessRuntimes/serverlessRuntimeName"),
    },
};
InformaticaOrganizationResource result = await informaticaOrganization.UpdateAsync(patch);

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
InformaticaOrganizationData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
