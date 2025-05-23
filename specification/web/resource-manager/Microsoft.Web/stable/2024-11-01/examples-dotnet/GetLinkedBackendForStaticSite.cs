using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.AppService;

// Generated from example definition: specification/web/resource-manager/Microsoft.Web/stable/2024-11-01/examples/GetLinkedBackendForStaticSite.json
// this example is just showing the usage of "StaticSites_GetLinkedBackend" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this StaticSiteLinkedBackendResource created on azure
// for more information of creating StaticSiteLinkedBackendResource, please refer to the document of StaticSiteLinkedBackendResource
string subscriptionId = "34adfa4f-cedf-4dc0-ba29-b6d1a69ab345";
string resourceGroupName = "rg";
string name = "testStaticSite0";
string linkedBackendName = "testBackend";
ResourceIdentifier staticSiteLinkedBackendResourceId = StaticSiteLinkedBackendResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, name, linkedBackendName);
StaticSiteLinkedBackendResource staticSiteLinkedBackend = client.GetStaticSiteLinkedBackendResource(staticSiteLinkedBackendResourceId);

// invoke the operation
StaticSiteLinkedBackendResource result = await staticSiteLinkedBackend.GetAsync();

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
StaticSiteLinkedBackendData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
