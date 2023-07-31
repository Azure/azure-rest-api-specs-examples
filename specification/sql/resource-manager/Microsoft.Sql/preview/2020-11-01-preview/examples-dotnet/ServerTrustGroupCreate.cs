using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Resources;
using Azure.ResourceManager.Sql;
using Azure.ResourceManager.Sql.Models;

// Generated from example definition: specification/sql/resource-manager/Microsoft.Sql/preview/2020-11-01-preview/examples/ServerTrustGroupCreate.json
// this example is just showing the usage of "ServerTrustGroups_CreateOrUpdate" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ResourceGroupResource created on azure
// for more information of creating ResourceGroupResource, please refer to the document of ResourceGroupResource
string subscriptionId = "00000000-1111-2222-3333-444444444444";
string resourceGroupName = "Default";
ResourceIdentifier resourceGroupResourceId = ResourceGroupResource.CreateResourceIdentifier(subscriptionId, resourceGroupName);
ResourceGroupResource resourceGroupResource = client.GetResourceGroupResource(resourceGroupResourceId);

// get the collection of this SqlServerTrustGroupResource
AzureLocation locationName = new AzureLocation("Japan East");
SqlServerTrustGroupCollection collection = resourceGroupResource.GetSqlServerTrustGroups(locationName);

// invoke the operation
string serverTrustGroupName = "server-trust-group-test";
SqlServerTrustGroupData data = new SqlServerTrustGroupData()
{
    GroupMembers =
    {
    new ServerTrustGroupServerInfo(new ResourceIdentifier("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/Default/providers/Microsoft.Sql/managedInstances/managedInstance-1")),new ServerTrustGroupServerInfo(new ResourceIdentifier("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/Default/providers/Microsoft.Sql/managedInstances/managedInstance-2"))
    },
    TrustScopes =
    {
    ServerTrustGroupPropertiesTrustScopesItem.GlobalTransactions,ServerTrustGroupPropertiesTrustScopesItem.ServiceBroker
    },
};
ArmOperation<SqlServerTrustGroupResource> lro = await collection.CreateOrUpdateAsync(WaitUntil.Completed, serverTrustGroupName, data);
SqlServerTrustGroupResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
SqlServerTrustGroupData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
