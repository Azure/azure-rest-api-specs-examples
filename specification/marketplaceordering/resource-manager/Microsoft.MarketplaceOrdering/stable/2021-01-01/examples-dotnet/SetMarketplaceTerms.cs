using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.MarketplaceOrdering.Models;
using Azure.ResourceManager.Resources;
using Azure.ResourceManager.MarketplaceOrdering;

// Generated from example definition: specification/marketplaceordering/resource-manager/Microsoft.MarketplaceOrdering/stable/2021-01-01/examples/SetMarketplaceTerms.json
// this example is just showing the usage of "MarketplaceAgreements_Create" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this SubscriptionResource created on azure
// for more information of creating SubscriptionResource, please refer to the document of SubscriptionResource
string subscriptionId = "subid";
ResourceIdentifier subscriptionResourceId = SubscriptionResource.CreateResourceIdentifier(subscriptionId);
SubscriptionResource subscriptionResource = client.GetSubscriptionResource(subscriptionResourceId);

// get the collection of this MarketplaceAgreementTermResource
MarketplaceAgreementTermCollection collection = subscriptionResource.GetMarketplaceAgreementTerms();

// invoke the operation
AgreementOfferType offerType = AgreementOfferType.Virtualmachine;
string publisherId = "pubid";
string offerId = "offid";
string planId = "planid";
MarketplaceAgreementTermData data = new MarketplaceAgreementTermData
{
    Publisher = "pubid",
    Product = "offid",
    Plan = "planid",
    LicenseTextLink = new Uri("test.licenseLink"),
    PrivacyPolicyLink = new Uri("test.privacyPolicyLink"),
    MarketplaceTermsLink = new Uri("test.marketplaceTermsLink"),
    RetrievedOn = DateTimeOffset.Parse("2017-08-15T11:33:07.12132Z"),
    Signature = "ASDFSDAFWEFASDGWERLWER",
    IsAccepted = false,
};
ArmOperation<MarketplaceAgreementTermResource> lro = await collection.CreateOrUpdateAsync(WaitUntil.Completed, offerType, publisherId, offerId, planId, data);
MarketplaceAgreementTermResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
MarketplaceAgreementTermData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
