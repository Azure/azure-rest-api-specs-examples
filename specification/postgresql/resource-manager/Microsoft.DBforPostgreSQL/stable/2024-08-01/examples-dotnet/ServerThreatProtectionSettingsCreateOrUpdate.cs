using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.PostgreSql.FlexibleServers.Models;
using Azure.ResourceManager.PostgreSql.FlexibleServers;

// Generated from example definition: specification/postgresql/resource-manager/Microsoft.DBforPostgreSQL/stable/2024-08-01/examples/ServerThreatProtectionSettingsCreateOrUpdate.json
// this example is just showing the usage of "ServerThreatProtectionSettings_CreateOrUpdate" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this PostgreSqlFlexibleServerResource created on azure
// for more information of creating PostgreSqlFlexibleServerResource, please refer to the document of PostgreSqlFlexibleServerResource
string subscriptionId = "00000000-1111-2222-3333-444444444444";
string resourceGroupName = "threatprotection-4799";
string serverName = "threatprotection-6440";
ResourceIdentifier postgreSqlFlexibleServerResourceId = PostgreSqlFlexibleServerResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, serverName);
PostgreSqlFlexibleServerResource postgreSqlFlexibleServer = client.GetPostgreSqlFlexibleServerResource(postgreSqlFlexibleServerResourceId);

// get the collection of this ServerThreatProtectionSettingsModelResource
ServerThreatProtectionSettingsModelCollection collection = postgreSqlFlexibleServer.GetServerThreatProtectionSettingsModels();

// invoke the operation
ThreatProtectionName threatProtectionName = ThreatProtectionName.Default;
ServerThreatProtectionSettingsModelData data = new ServerThreatProtectionSettingsModelData
{
    State = ThreatProtectionState.Enabled,
};
ArmOperation<ServerThreatProtectionSettingsModelResource> lro = await collection.CreateOrUpdateAsync(WaitUntil.Completed, threatProtectionName, data);
ServerThreatProtectionSettingsModelResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
ServerThreatProtectionSettingsModelData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
