using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.RecoveryServicesSiteRecovery.Models;
using Azure.ResourceManager.RecoveryServicesSiteRecovery;

// Generated from example definition: specification/recoveryservicessiterecovery/resource-manager/Microsoft.RecoveryServices/stable/2025-01-01/examples/ReplicationAlertSettings_Create.json
// this example is just showing the usage of "ReplicationAlertSettings_Create" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this SiteRecoveryAlertResource created on azure
// for more information of creating SiteRecoveryAlertResource, please refer to the document of SiteRecoveryAlertResource
string subscriptionId = "c183865e-6077-46f2-a3b1-deb0f4f4650a";
string resourceGroupName = "resourceGroupPS1";
string resourceName = "vault1";
string alertSettingName = "defaultAlertSetting";
ResourceIdentifier siteRecoveryAlertResourceId = SiteRecoveryAlertResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, resourceName, alertSettingName);
SiteRecoveryAlertResource siteRecoveryAlert = client.GetSiteRecoveryAlertResource(siteRecoveryAlertResourceId);

// invoke the operation
SiteRecoveryAlertCreateOrUpdateContent content = new SiteRecoveryAlertCreateOrUpdateContent
{
    Properties = new SiteRecoveryConfigureAlertProperties
    {
        SendToOwners = "false",
        CustomEmailAddresses = { "ronehr@microsoft.com" },
        Locale = "",
    },
};
ArmOperation<SiteRecoveryAlertResource> lro = await siteRecoveryAlert.UpdateAsync(WaitUntil.Completed, content);
SiteRecoveryAlertResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
SiteRecoveryAlertData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
