using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.AppService;

// Generated from example definition: specification/web/resource-manager/Microsoft.Web/stable/2021-02-01/examples/GetUserProvidedFunctionAppForStaticSiteBuild.json
// this example is just showing the usage of "StaticSites_GetUserProvidedFunctionAppForStaticSiteBuild" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this StaticSiteBuildResource created on azure
// for more information of creating StaticSiteBuildResource, please refer to the document of StaticSiteBuildResource
string subscriptionId = "34adfa4f-cedf-4dc0-ba29-b6d1a69ab345";
string resourceGroupName = "rg";
string name = "testStaticSite0";
string environmentName = "default";
ResourceIdentifier staticSiteBuildResourceId = StaticSiteBuildResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, name, environmentName);
StaticSiteBuildResource staticSiteBuild = client.GetStaticSiteBuildResource(staticSiteBuildResourceId);

// get the collection of this StaticSiteBuildUserProvidedFunctionAppResource
StaticSiteBuildUserProvidedFunctionAppCollection collection = staticSiteBuild.GetStaticSiteBuildUserProvidedFunctionApps();

// invoke the operation
string functionAppName = "testFunctionApp";
bool result = await collection.ExistsAsync(functionAppName);

Console.WriteLine($"Succeeded: {result}");
