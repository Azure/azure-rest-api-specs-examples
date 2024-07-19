using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.AppPlatform.Models;
using Azure.ResourceManager.AppPlatform;

// Generated from example definition: specification/appplatform/resource-manager/Microsoft.AppPlatform/stable/2022-12-01/examples/ConfigurationServices_Validate.json
// this example is just showing the usage of "ConfigurationServices_Validate" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this AppPlatformConfigurationServiceResource created on azure
// for more information of creating AppPlatformConfigurationServiceResource, please refer to the document of AppPlatformConfigurationServiceResource
string subscriptionId = "00000000-0000-0000-0000-000000000000";
string resourceGroupName = "myResourceGroup";
string serviceName = "myservice";
string configurationServiceName = "default";
ResourceIdentifier appPlatformConfigurationServiceResourceId = AppPlatformConfigurationServiceResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, serviceName, configurationServiceName);
AppPlatformConfigurationServiceResource appPlatformConfigurationService = client.GetAppPlatformConfigurationServiceResource(appPlatformConfigurationServiceResourceId);

// invoke the operation
AppPlatformConfigurationServiceSettings settings = new AppPlatformConfigurationServiceSettings()
{
    ConfigurationServiceGitRepositories =
    {
    new AppPlatformConfigurationServiceGitRepository("fake",new string[]
    {
    "app/dev"
    },new Uri("https://github.com/fake-user/fake-repository"),"master")
    },
};
ArmOperation<AppPlatformConfigurationServiceSettingsValidateResult> lro = await appPlatformConfigurationService.ValidateAsync(WaitUntil.Completed, settings);
AppPlatformConfigurationServiceSettingsValidateResult result = lro.Value;

Console.WriteLine($"Succeeded: {result}");
