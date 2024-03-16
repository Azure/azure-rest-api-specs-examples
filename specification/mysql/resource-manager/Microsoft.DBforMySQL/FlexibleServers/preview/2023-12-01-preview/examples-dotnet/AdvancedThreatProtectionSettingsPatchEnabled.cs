using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.MySql.FlexibleServers;
using Azure.ResourceManager.MySql.FlexibleServers.Models;

// Generated from example definition: specification/mysql/resource-manager/Microsoft.DBforMySQL/FlexibleServers/preview/2023-12-01-preview/examples/AdvancedThreatProtectionSettingsPatchEnabled.json
// this example is just showing the usage of "AdvancedThreatProtectionSettings_Update" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this AdvancedThreatProtectionResource created on azure
// for more information of creating AdvancedThreatProtectionResource, please refer to the document of AdvancedThreatProtectionResource
string subscriptionId = "00000000-1111-2222-3333-444444444444";
string resourceGroupName = "threatprotection-4799";
string serverName = "threatprotection-6440";
AdvancedThreatProtectionName advancedThreatProtectionName = AdvancedThreatProtectionName.Default;
ResourceIdentifier advancedThreatProtectionResourceId = AdvancedThreatProtectionResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, serverName, advancedThreatProtectionName);
AdvancedThreatProtectionResource advancedThreatProtection = client.GetAdvancedThreatProtectionResource(advancedThreatProtectionResourceId);

// invoke the operation
AdvancedThreatProtectionPatch patch = new AdvancedThreatProtectionPatch()
{
    State = AdvancedThreatProtectionState.Enabled,
};
ArmOperation<AdvancedThreatProtectionResource> lro = await advancedThreatProtection.UpdateAsync(WaitUntil.Completed, patch);
AdvancedThreatProtectionResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
AdvancedThreatProtectionData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
