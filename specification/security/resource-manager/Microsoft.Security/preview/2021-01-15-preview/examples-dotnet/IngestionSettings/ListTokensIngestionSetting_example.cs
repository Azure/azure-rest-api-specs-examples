using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.SecurityCenter;
using Azure.ResourceManager.SecurityCenter.Models;

// Generated from example definition: specification/security/resource-manager/Microsoft.Security/preview/2021-01-15-preview/examples/IngestionSettings/ListTokensIngestionSetting_example.json
// this example is just showing the usage of "IngestionSettings_ListTokens" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this IngestionSettingResource created on azure
// for more information of creating IngestionSettingResource, please refer to the document of IngestionSettingResource
string subscriptionId = "20ff7fc3-e762-44dd-bd96-b71116dcdc23";
string ingestionSettingName = "default";
ResourceIdentifier ingestionSettingResourceId = IngestionSettingResource.CreateResourceIdentifier(subscriptionId, ingestionSettingName);
IngestionSettingResource ingestionSetting = client.GetIngestionSettingResource(ingestionSettingResourceId);

// invoke the operation
IngestionSettingToken result = await ingestionSetting.GetTokensAsync();

Console.WriteLine($"Succeeded: {result}");
