using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Resources;
using Azure.ResourceManager.SiteManager.Models;
using Azure.ResourceManager.SiteManager;

// Generated from example definition: 2025-03-01-preview/SitesBySubscription_CreateOrUpdate_MaximumSet_Gen.json
// this example is just showing the usage of "Site_CreateOrUpdate" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this SubscriptionResource created on azure
// for more information of creating SubscriptionResource, please refer to the document of SubscriptionResource
string subscriptionId = "0154f7fe-df09-4981-bf82-7ad5c1f596eb";
ResourceIdentifier subscriptionResourceId = SubscriptionResource.CreateResourceIdentifier(subscriptionId);
SubscriptionResource subscriptionResource = client.GetSubscriptionResource(subscriptionResourceId);

// invoke the operation
string siteName = "string";
EdgeSiteData data = new EdgeSiteData
{
    Properties = new EdgeSiteProperties
    {
        DisplayName = "string",
        Description = "enxcmpvfvadbapo",
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
        ["key8188"] = "mcgnu"
        },
    },
};
ArmOperation<EdgeSiteResource> lro = await subscriptionResource.CreateOrUpdateSitesBySubscriptionAsync(WaitUntil.Completed, siteName, data);
EdgeSiteResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
EdgeSiteData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
