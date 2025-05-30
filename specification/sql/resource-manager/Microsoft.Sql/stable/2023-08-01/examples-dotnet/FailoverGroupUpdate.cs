using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Sql.Models;
using Azure.ResourceManager.Sql;

// Generated from example definition: specification/sql/resource-manager/Microsoft.Sql/stable/2023-08-01/examples/FailoverGroupUpdate.json
// this example is just showing the usage of "FailoverGroups_Update" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this FailoverGroupResource created on azure
// for more information of creating FailoverGroupResource, please refer to the document of FailoverGroupResource
string subscriptionId = "00000000-1111-2222-3333-444444444444";
string resourceGroupName = "Default";
string serverName = "failover-group-primary-server";
string failoverGroupName = "failover-group-test-1";
ResourceIdentifier failoverGroupResourceId = FailoverGroupResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, serverName, failoverGroupName);
FailoverGroupResource failoverGroup = client.GetFailoverGroupResource(failoverGroupResourceId);

// invoke the operation
FailoverGroupPatch patch = new FailoverGroupPatch
{
    ReadWriteEndpoint = new FailoverGroupReadWriteEndpoint(ReadWriteEndpointFailoverPolicy.Automatic)
    {
        FailoverWithDataLossGracePeriodMinutes = 120,
    },
    FailoverDatabases = { new ResourceIdentifier("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/Default/providers/Microsoft.Sql/servers/failover-group-primary-server/databases/testdb-1") },
};
ArmOperation<FailoverGroupResource> lro = await failoverGroup.UpdateAsync(WaitUntil.Completed, patch);
FailoverGroupResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
FailoverGroupData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
