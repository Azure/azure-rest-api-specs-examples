using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.AppService.Models;
using Azure.ResourceManager.AppService;

// Generated from example definition: specification/web/resource-manager/Microsoft.Web/stable/2024-11-01/examples/DeleteSiteCertificateSlot.json
// this example is just showing the usage of "SiteCertificates_DeleteSlot" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this SiteSlotCertificateResource created on azure
// for more information of creating SiteSlotCertificateResource, please refer to the document of SiteSlotCertificateResource
string subscriptionId = "34adfa4f-cedf-4dc0-ba29-b6d1a69ab345";
string resourceGroupName = "testrg123";
string name = "testSiteName";
string slot = "staging";
string certificateName = "testc6282";
ResourceIdentifier siteSlotCertificateResourceId = SiteSlotCertificateResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, name, slot, certificateName);
SiteSlotCertificateResource siteSlotCertificate = client.GetSiteSlotCertificateResource(siteSlotCertificateResourceId);

// invoke the operation
await siteSlotCertificate.DeleteAsync(WaitUntil.Completed);

Console.WriteLine("Succeeded");
