using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.SecurityCenter.Models;
using Azure.ResourceManager.SecurityCenter;

// Generated from example definition: specification/security/resource-manager/Microsoft.Security/preview/2023-02-15-preview/examples/SensitivitySettings/GetSensitivitySettings_example.json
// this example is just showing the usage of "GetSensitivitySettings" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this SensitivitySettingResource created on azure
// for more information of creating SensitivitySettingResource, please refer to the document of SensitivitySettingResource
ResourceIdentifier sensitivitySettingResourceId = SensitivitySettingResource.CreateResourceIdentifier();
SensitivitySettingResource sensitivitySetting = client.GetSensitivitySettingResource(sensitivitySettingResourceId);

// invoke the operation
SensitivitySettingResource result = await sensitivitySetting.GetAsync();

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
SensitivitySettingData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
