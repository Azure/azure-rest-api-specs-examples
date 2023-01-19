using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.AppService;
using Azure.ResourceManager.AppService.Models;

// Generated from example definition: specification/web/resource-manager/Microsoft.Web/stable/2021-02-01/examples/CreateOrUpdateStaticSiteBuildFunctionAppSettings.json
// this example is just showing the usage of "StaticSites_CreateOrUpdateStaticSiteBuildFunctionAppSettings" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this StaticSiteBuildResource created on azure
// for more information of creating StaticSiteBuildResource, please refer to the document of StaticSiteBuildResource
string subscriptionId = "34adfa4f-cedf-4dc0-ba29-b6d1a69ab345";
string resourceGroupName = "rg";
string name = "testStaticSite0";
string environmentName = "12";
ResourceIdentifier staticSiteBuildResourceId = StaticSiteBuildResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, name, environmentName);
StaticSiteBuildResource staticSiteBuild = client.GetStaticSiteBuildResource(staticSiteBuildResourceId);

// invoke the operation
AppServiceConfigurationDictionary appSettings = new AppServiceConfigurationDictionary()
{
    Properties =
    {
    ["setting1"] = "someval",
    ["setting2"] = "someval2",
    },
};
AppServiceConfigurationDictionary result = await staticSiteBuild.CreateOrUpdateFunctionAppSettingsAsync(appSettings);

Console.WriteLine($"Succeeded: {result}");
