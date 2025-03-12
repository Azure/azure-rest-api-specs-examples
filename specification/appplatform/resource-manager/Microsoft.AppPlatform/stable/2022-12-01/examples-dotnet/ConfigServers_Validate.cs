using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.AppPlatform.Models;
using Azure.ResourceManager.AppPlatform;

// Generated from example definition: specification/appplatform/resource-manager/Microsoft.AppPlatform/stable/2022-12-01/examples/ConfigServers_Validate.json
// this example is just showing the usage of "ConfigServers_Validate" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this AppPlatformServiceResource created on azure
// for more information of creating AppPlatformServiceResource, please refer to the document of AppPlatformServiceResource
string subscriptionId = "00000000-0000-0000-0000-000000000000";
string resourceGroupName = "myResourceGroup";
string serviceName = "myservice";
ResourceIdentifier appPlatformServiceResourceId = AppPlatformServiceResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, serviceName);
AppPlatformServiceResource appPlatformService = client.GetAppPlatformServiceResource(appPlatformServiceResourceId);

// invoke the operation
ConfigServerSettings settings = new ConfigServerSettings
{
    GitProperty = new AppPlatformConfigServerGitProperty(new Uri("https://github.com/fake-user/fake-repository.git"))
    {
        Label = "master",
        SearchPaths = { "/" },
    },
};
ArmOperation<ConfigServerSettingsValidateResult> lro = await appPlatformService.ValidateConfigServerAsync(WaitUntil.Completed, settings);
ConfigServerSettingsValidateResult result = lro.Value;

Console.WriteLine($"Succeeded: {result}");
