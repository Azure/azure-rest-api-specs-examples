using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.CloudHealth.Models;
using Azure.ResourceManager.CloudHealth;

// Generated from example definition: 2025-05-01-preview/AuthenticationSettings_Delete.json
// this example is just showing the usage of "AuthenticationSetting_Delete" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this HealthModelAuthenticationSettingResource created on azure
// for more information of creating HealthModelAuthenticationSettingResource, please refer to the document of HealthModelAuthenticationSettingResource
string subscriptionId = "00000000-0000-0000-0000-000000000000";
string resourceGroupName = "my-resource-group";
string healthModelName = "my-health-model";
string authenticationSettingName = "my-auth-setting";
ResourceIdentifier healthModelAuthenticationSettingResourceId = HealthModelAuthenticationSettingResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, healthModelName, authenticationSettingName);
HealthModelAuthenticationSettingResource healthModelAuthenticationSetting = client.GetHealthModelAuthenticationSettingResource(healthModelAuthenticationSettingResourceId);

// invoke the operation
await healthModelAuthenticationSetting.DeleteAsync(WaitUntil.Completed);

Console.WriteLine("Succeeded");
