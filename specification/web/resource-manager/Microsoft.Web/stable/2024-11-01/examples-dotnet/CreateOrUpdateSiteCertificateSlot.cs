using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.AppService;

// Generated from example definition: specification/web/resource-manager/Microsoft.Web/stable/2024-11-01/examples/CreateOrUpdateSiteCertificateSlot.json
// this example is just showing the usage of "SiteCertificates_CreateOrUpdateSlot" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this WebSiteSlotResource created on azure
// for more information of creating WebSiteSlotResource, please refer to the document of WebSiteSlotResource
string subscriptionId = "34adfa4f-cedf-4dc0-ba29-b6d1a69ab345";
string resourceGroupName = "testrg123";
string name = "testSiteName";
string slot = "staging";
ResourceIdentifier webSiteSlotResourceId = WebSiteSlotResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, name, slot);
WebSiteSlotResource webSiteSlot = client.GetWebSiteSlotResource(webSiteSlotResourceId);

// get the collection of this SiteSlotCertificateResource
SiteSlotCertificateCollection collection = webSiteSlot.GetSiteSlotCertificates();

// invoke the operation
string certificateName = "testc6282";
AppCertificateData data = new AppCertificateData(new AzureLocation("East US"))
{
    Password = "<password>",
    HostNames = { "ServerCert" },
};
ArmOperation<SiteSlotCertificateResource> lro = await collection.CreateOrUpdateAsync(WaitUntil.Completed, certificateName, data);
SiteSlotCertificateResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
AppCertificateData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
