using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.MarketplaceOrdering;

// Generated from example definition: specification/marketplaceordering/resource-manager/Microsoft.MarketplaceOrdering/stable/2021-01-01/examples/CancelMarketplaceTerms.json
// this example is just showing the usage of "MarketplaceAgreements_Cancel" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this MarketplaceAgreementResource created on azure
// for more information of creating MarketplaceAgreementResource, please refer to the document of MarketplaceAgreementResource
string subscriptionId = "subid";
string publisherId = "pubid";
string offerId = "offid";
string planId = "planid";
ResourceIdentifier marketplaceAgreementResourceId = MarketplaceAgreementResource.CreateResourceIdentifier(subscriptionId, publisherId, offerId, planId);
MarketplaceAgreementResource marketplaceAgreement = client.GetMarketplaceAgreementResource(marketplaceAgreementResourceId);

// invoke the operation
MarketplaceAgreementResource result = await marketplaceAgreement.CancelAsync();

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
MarketplaceAgreementTermData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
