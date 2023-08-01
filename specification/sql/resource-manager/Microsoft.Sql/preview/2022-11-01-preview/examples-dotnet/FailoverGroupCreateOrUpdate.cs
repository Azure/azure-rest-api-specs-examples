using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Sql;
using Azure.ResourceManager.Sql.Models;

// Generated from example definition: specification/sql/resource-manager/Microsoft.Sql/preview/2022-11-01-preview/examples/FailoverGroupCreateOrUpdate.json
// this example is just showing the usage of "FailoverGroups_CreateOrUpdate" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this SqlServerResource created on azure
// for more information of creating SqlServerResource, please refer to the document of SqlServerResource
string subscriptionId = "00000000-1111-2222-3333-444444444444";
string resourceGroupName = "Default";
string serverName = "failover-group-primary-server";
ResourceIdentifier sqlServerResourceId = SqlServerResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, serverName);
SqlServerResource sqlServer = client.GetSqlServerResource(sqlServerResourceId);

// get the collection of this FailoverGroupResource
FailoverGroupCollection collection = sqlServer.GetFailoverGroups();

// invoke the operation
string failoverGroupName = "failover-group-test-3";
FailoverGroupData data = new FailoverGroupData()
{
    ReadWriteEndpoint = new FailoverGroupReadWriteEndpoint(ReadWriteEndpointFailoverPolicy.Automatic)
    {
        FailoverWithDataLossGracePeriodMinutes = 480,
    },
    ReadOnlyEndpointFailoverPolicy = ReadOnlyEndpointFailoverPolicy.Disabled,
    PartnerServers =
    {
    new PartnerServerInfo(new ResourceIdentifier("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/Default/providers/Microsoft.Sql/servers/failover-group-secondary-server"))
    },
    FailoverDatabases =
    {
    new ResourceIdentifier("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/Default/providers/Microsoft.Sql/servers/failover-group-primary-server/databases/testdb-1"),new ResourceIdentifier("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/Default/providers/Microsoft.Sql/servers/failover-group-primary-server/databases/testdb-2")
    },
};
ArmOperation<FailoverGroupResource> lro = await collection.CreateOrUpdateAsync(WaitUntil.Completed, failoverGroupName, data);
FailoverGroupResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
FailoverGroupData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
