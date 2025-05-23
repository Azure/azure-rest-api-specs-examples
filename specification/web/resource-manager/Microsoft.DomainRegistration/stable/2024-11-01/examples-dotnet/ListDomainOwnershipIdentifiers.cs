using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.AppService;

// Generated from example definition: specification/web/resource-manager/Microsoft.DomainRegistration/stable/2024-11-01/examples/ListDomainOwnershipIdentifiers.json
// this example is just showing the usage of "Domains_ListOwnershipIdentifiers" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this AppServiceDomainResource created on azure
// for more information of creating AppServiceDomainResource, please refer to the document of AppServiceDomainResource
string subscriptionId = "34adfa4f-cedf-4dc0-ba29-b6d1a69ab345";
string resourceGroupName = "testrg123";
string domainName = "example.com";
ResourceIdentifier appServiceDomainResourceId = AppServiceDomainResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, domainName);
AppServiceDomainResource appServiceDomain = client.GetAppServiceDomainResource(appServiceDomainResourceId);

// get the collection of this DomainOwnershipIdentifierResource
DomainOwnershipIdentifierCollection collection = appServiceDomain.GetDomainOwnershipIdentifiers();

// invoke the operation and iterate over the result
await foreach (DomainOwnershipIdentifierResource item in collection.GetAllAsync())
{
    // the variable item is a resource, you could call other operations on this instance as well
    // but just for demo, we get its data from this resource instance
    DomainOwnershipIdentifierData resourceData = item.Data;
    // for demo we just print out the id
    Console.WriteLine($"Succeeded on id: {resourceData.Id}");
}

Console.WriteLine("Succeeded");
