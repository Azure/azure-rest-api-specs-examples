using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.SecurityCenter;
using Azure.ResourceManager.SecurityCenter.Models;

// Generated from example definition: specification/security/resource-manager/Microsoft.Security/stable/2022-05-01/examples/Settings/GetSetting_example.json
// this example is just showing the usage of "Settings_Get" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this SecuritySettingResource created on azure
// for more information of creating SecuritySettingResource, please refer to the document of SecuritySettingResource
string subscriptionId = "20ff7fc3-e762-44dd-bd96-b71116dcdc23";
SecuritySettingName settingName = SecuritySettingName.Wdatp;
ResourceIdentifier securitySettingResourceId = SecuritySettingResource.CreateResourceIdentifier(subscriptionId, settingName);
SecuritySettingResource securitySetting = client.GetSecuritySettingResource(securitySettingResourceId);

// invoke the operation
SecuritySettingResource result = await securitySetting.GetAsync();

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
SecuritySettingData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
