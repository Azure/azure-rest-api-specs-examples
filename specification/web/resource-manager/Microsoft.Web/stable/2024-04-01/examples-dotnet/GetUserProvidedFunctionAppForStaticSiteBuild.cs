using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.AppService;

// Generated from example definition: specification/web/resource-manager/Microsoft.Web/stable/2024-04-01/examples/GetUserProvidedFunctionAppForStaticSiteBuild.json
// this example is just showing the usage of "StaticSites_GetUserProvidedFunctionAppForStaticSiteBuild" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this StaticSiteBuildUserProvidedFunctionAppResource created on azure
// for more information of creating StaticSiteBuildUserProvidedFunctionAppResource, please refer to the document of StaticSiteBuildUserProvidedFunctionAppResource
string subscriptionId = "34adfa4f-cedf-4dc0-ba29-b6d1a69ab345";
string resourceGroupName = "rg";
string name = "testStaticSite0";
string environmentName = "default";
string functionAppName = "testFunctionApp";
ResourceIdentifier staticSiteBuildUserProvidedFunctionAppResourceId = StaticSiteBuildUserProvidedFunctionAppResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, name, environmentName, functionAppName);
StaticSiteBuildUserProvidedFunctionAppResource staticSiteBuildUserProvidedFunctionApp = client.GetStaticSiteBuildUserProvidedFunctionAppResource(staticSiteBuildUserProvidedFunctionAppResourceId);

// invoke the operation
StaticSiteBuildUserProvidedFunctionAppResource result = await staticSiteBuildUserProvidedFunctionApp.GetAsync();

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
StaticSiteUserProvidedFunctionAppData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
