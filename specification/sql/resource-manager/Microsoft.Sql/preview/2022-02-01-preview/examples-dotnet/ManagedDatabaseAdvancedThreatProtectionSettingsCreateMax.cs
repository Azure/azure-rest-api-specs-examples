using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Sql;
using Azure.ResourceManager.Sql.Models;

// Generated from example definition: specification/sql/resource-manager/Microsoft.Sql/preview/2022-02-01-preview/examples/ManagedDatabaseAdvancedThreatProtectionSettingsCreateMax.json
// this example is just showing the usage of "ManagedDatabaseAdvancedThreatProtectionSettings_CreateOrUpdate" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ManagedDatabaseAdvancedThreatProtectionResource created on azure
// for more information of creating ManagedDatabaseAdvancedThreatProtectionResource, please refer to the document of ManagedDatabaseAdvancedThreatProtectionResource
string subscriptionId = "00000000-1111-2222-3333-444444444444";
string resourceGroupName = "threatprotection-4799";
string managedInstanceName = "threatprotection-6440";
string databaseName = "testdb";
AdvancedThreatProtectionName advancedThreatProtectionName = AdvancedThreatProtectionName.Default;
ResourceIdentifier managedDatabaseAdvancedThreatProtectionResourceId = ManagedDatabaseAdvancedThreatProtectionResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, managedInstanceName, databaseName, advancedThreatProtectionName);
ManagedDatabaseAdvancedThreatProtectionResource managedDatabaseAdvancedThreatProtection = client.GetManagedDatabaseAdvancedThreatProtectionResource(managedDatabaseAdvancedThreatProtectionResourceId);

// invoke the operation
ManagedDatabaseAdvancedThreatProtectionData data = new ManagedDatabaseAdvancedThreatProtectionData()
{
    State = AdvancedThreatProtectionState.Enabled,
};
ArmOperation<ManagedDatabaseAdvancedThreatProtectionResource> lro = await managedDatabaseAdvancedThreatProtection.UpdateAsync(WaitUntil.Completed, data);
ManagedDatabaseAdvancedThreatProtectionResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
ManagedDatabaseAdvancedThreatProtectionData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
