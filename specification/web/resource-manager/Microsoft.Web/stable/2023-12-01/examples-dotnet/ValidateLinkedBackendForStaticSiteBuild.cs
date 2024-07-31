using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.AppService;

// Generated from example definition: specification/web/resource-manager/Microsoft.Web/stable/2023-12-01/examples/ValidateLinkedBackendForStaticSiteBuild.json
// this example is just showing the usage of "StaticSites_ValidateBackendForBuild" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this StaticSiteBuildLinkedBackendResource created on azure
// for more information of creating StaticSiteBuildLinkedBackendResource, please refer to the document of StaticSiteBuildLinkedBackendResource
string subscriptionId = "34adfa4f-cedf-4dc0-ba29-b6d1a69ab345";
string resourceGroupName = "rg";
string name = "testStaticSite0";
string environmentName = "default";
string linkedBackendName = "testBackend";
ResourceIdentifier staticSiteBuildLinkedBackendResourceId = StaticSiteBuildLinkedBackendResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, name, environmentName, linkedBackendName);
StaticSiteBuildLinkedBackendResource staticSiteBuildLinkedBackend = client.GetStaticSiteBuildLinkedBackendResource(staticSiteBuildLinkedBackendResourceId);

// invoke the operation
StaticSiteLinkedBackendData data = new StaticSiteLinkedBackendData()
{
    BackendResourceId = new ResourceIdentifier("/subscription/34adfa4f-cedf-4dc0-ba29-b6d1a69ab345/resourceGroups/backendRg/providers/Microsoft.Web/sites/testBackend"),
    Region = "West US 2",
};
await staticSiteBuildLinkedBackend.ValidateBackendForBuildAsync(WaitUntil.Completed, data);

Console.WriteLine($"Succeeded");
