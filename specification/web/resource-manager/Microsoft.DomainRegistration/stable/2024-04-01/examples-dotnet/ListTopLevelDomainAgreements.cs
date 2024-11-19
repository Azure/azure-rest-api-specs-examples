using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.AppService.Models;
using Azure.ResourceManager.AppService;

// Generated from example definition: specification/web/resource-manager/Microsoft.DomainRegistration/stable/2024-04-01/examples/ListTopLevelDomainAgreements.json
// this example is just showing the usage of "TopLevelDomains_ListAgreements" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this TopLevelDomainResource created on azure
// for more information of creating TopLevelDomainResource, please refer to the document of TopLevelDomainResource
string subscriptionId = "34adfa4f-cedf-4dc0-ba29-b6d1a69ab345";
string name = "in";
ResourceIdentifier topLevelDomainResourceId = TopLevelDomainResource.CreateResourceIdentifier(subscriptionId, name);
TopLevelDomainResource topLevelDomain = client.GetTopLevelDomainResource(topLevelDomainResourceId);

// invoke the operation and iterate over the result
TopLevelDomainAgreementOption agreementOption = new TopLevelDomainAgreementOption()
{
    IncludePrivacy = true,
    IsForTransfer = false,
};
await foreach (TldLegalAgreement item in topLevelDomain.GetAgreementsAsync(agreementOption))
{
    Console.WriteLine($"Succeeded: {item}");
}

Console.WriteLine("Succeeded");
