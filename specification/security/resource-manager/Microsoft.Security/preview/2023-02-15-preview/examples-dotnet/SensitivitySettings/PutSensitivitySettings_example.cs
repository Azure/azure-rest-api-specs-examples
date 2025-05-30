using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.SecurityCenter.Models;
using Azure.ResourceManager.SecurityCenter;

// Generated from example definition: specification/security/resource-manager/Microsoft.Security/preview/2023-02-15-preview/examples/SensitivitySettings/PutSensitivitySettings_example.json
// this example is just showing the usage of "UpdateSensitivitySettings" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this SensitivitySettingResource created on azure
// for more information of creating SensitivitySettingResource, please refer to the document of SensitivitySettingResource
ResourceIdentifier sensitivitySettingResourceId = SensitivitySettingResource.CreateResourceIdentifier();
SensitivitySettingResource sensitivitySetting = client.GetSensitivitySettingResource(sensitivitySettingResourceId);

// invoke the operation
SensitivitySettingCreateOrUpdateContent content = new SensitivitySettingCreateOrUpdateContent(new Guid[] { Guid.Parse("f2f8a7a1-28c0-404b-9ab4-30a0a7af18cb"), Guid.Parse("b452f22b-f87d-4f48-8490-ecf0873325b5"), Guid.Parse("d59ee8b6-2618-404b-a5e7-aa377cd67543") })
{
    SensitivityThresholdLabelOrder = 2,
    SensitivityThresholdLabelId = Guid.Parse("f2f8a7a1-28c0-404b-9ab4-30a0a7af18cb"),
};
ArmOperation<SensitivitySettingResource> lro = await sensitivitySetting.CreateOrUpdateAsync(WaitUntil.Completed, content);
SensitivitySettingResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
SensitivitySettingData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
