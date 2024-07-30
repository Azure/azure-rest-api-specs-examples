using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.AppService.Models;
using Azure.ResourceManager.Resources;
using Azure.ResourceManager.AppService;

// Generated from example definition: specification/web/resource-manager/Microsoft.DomainRegistration/stable/2023-12-01/examples/DeleteAppServiceDomain.json
// this example is just showing the usage of "Domains_Delete" operation, for the dependent resources, they will have to be created separately.

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

// invoke the operation
bool? forceHardDeleteDomain = true;
await appServiceDomain.DeleteAsync(WaitUntil.Completed, forceHardDeleteDomain: forceHardDeleteDomain);

Console.WriteLine($"Succeeded");
