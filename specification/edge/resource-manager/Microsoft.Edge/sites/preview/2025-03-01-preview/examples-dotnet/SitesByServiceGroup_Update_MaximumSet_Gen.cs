using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Resources;
using Azure.ResourceManager.SiteManager.Models;
using Azure.ResourceManager.SiteManager;

// Generated from example definition: 2025-03-01-preview/SitesByServiceGroup_Update_MaximumSet_Gen.json
// this example is just showing the usage of "SitesByServiceGroup_Update" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

TenantResource tenantResource = client.GetTenants().GetAllAsync().GetAsyncEnumerator().Current;

// invoke the operation
string servicegroupName = "string";
string siteName = "string";
EdgeSitePatch patch = new EdgeSitePatch
{
    Properties = new EdgeSitePatchProperties
    {
        DisplayName = "string",
        Description = "zztr",
        SiteAddress = new SiteAddressProperties
        {
            StreetAddress1 = "fodimymrxbhrfslsmzfhmitn",
            StreetAddress2 = "widjg",
            City = "zkcbzjkatafo",
            StateOrProvince = "wk",
            Country = "xeevcfvimlfzsfuxtyujw",
            PostalCode = "qbrhqk",
        },
        Labels =
        {
        ["key9939"] = "jdlzxcvcfqmruq"
        },
    },
};
EdgeSiteResource result = await tenantResource.UpdateSitesByServiceGroupAsync(servicegroupName, siteName, patch);

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
EdgeSiteData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
