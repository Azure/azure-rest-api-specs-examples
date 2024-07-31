using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.AppService;

// Generated from example definition: specification/web/resource-manager/Microsoft.DomainRegistration/stable/2023-12-01/examples/DeleteAppServiceDomainOwnershipIdentifier.json
// this example is just showing the usage of "Domains_DeleteOwnershipIdentifier" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this DomainOwnershipIdentifierResource created on azure
// for more information of creating DomainOwnershipIdentifierResource, please refer to the document of DomainOwnershipIdentifierResource
string subscriptionId = "34adfa4f-cedf-4dc0-ba29-b6d1a69ab345";
string resourceGroupName = "testrg123";
string domainName = "example.com";
string name = "ownershipIdentifier";
ResourceIdentifier domainOwnershipIdentifierResourceId = DomainOwnershipIdentifierResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, domainName, name);
DomainOwnershipIdentifierResource domainOwnershipIdentifier = client.GetDomainOwnershipIdentifierResource(domainOwnershipIdentifierResourceId);

// invoke the operation
await domainOwnershipIdentifier.DeleteAsync(WaitUntil.Completed);

Console.WriteLine($"Succeeded");
