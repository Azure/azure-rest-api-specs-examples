using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.AppService.Models;
using Azure.ResourceManager.AppService;

// Generated from example definition: specification/web/resource-manager/Microsoft.Web/stable/2024-11-01/examples/PatchSiteCertificateSlot.json
// this example is just showing the usage of "SiteCertificates_UpdateSlot" operation, for the dependent resources, they will have to be created separately.

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
AppCertificatePatch patch = new AppCertificatePatch
{
    KeyVaultId = new ResourceIdentifier("/subscriptions/34adfa4f-cedf-4dc0-ba29-b6d1a69ab345/resourceGroups/testrg123/providers/Microsoft.KeyVault/vaults/testKV"),
};
SiteSlotCertificateResource result = await siteSlotCertificate.UpdateAsync(patch);

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
AppCertificateData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
