using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Resources;
using Azure.ResourceManager.Sql;
using Azure.ResourceManager.Sql.Models;

// Generated from example definition: specification/sql/resource-manager/Microsoft.Sql/preview/2022-08-01-preview/examples/ServerCreate.json
// this example is just showing the usage of "Servers_CreateOrUpdate" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ResourceGroupResource created on azure
// for more information of creating ResourceGroupResource, please refer to the document of ResourceGroupResource
string subscriptionId = "00000000-1111-2222-3333-444444444444";
string resourceGroupName = "sqlcrudtest-7398";
ResourceIdentifier resourceGroupResourceId = ResourceGroupResource.CreateResourceIdentifier(subscriptionId, resourceGroupName);
ResourceGroupResource resourceGroupResource = client.GetResourceGroupResource(resourceGroupResourceId);

// get the collection of this SqlServerResource
SqlServerCollection collection = resourceGroupResource.GetSqlServers();

// invoke the operation
string serverName = "sqlcrudtest-4645";
SqlServerData data = new SqlServerData(new AzureLocation("Japan East"))
{
    AdministratorLogin = "dummylogin",
    AdministratorLoginPassword = "PLACEHOLDER",
    PublicNetworkAccess = ServerNetworkAccessFlag.Enabled,
    Administrators = new ServerExternalAdministrator()
    {
        PrincipalType = SqlServerPrincipalType.User,
        Login = "bob@contoso.com",
        Sid = Guid.Parse("00000011-1111-2222-2222-123456789111"),
        TenantId = Guid.Parse("00000011-1111-2222-2222-123456789111"),
        IsAzureADOnlyAuthenticationEnabled = true,
    },
    RestrictOutboundNetworkAccess = ServerNetworkAccessFlag.Enabled,
};
ArmOperation<SqlServerResource> lro = await collection.CreateOrUpdateAsync(WaitUntil.Completed, serverName, data);
SqlServerResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
SqlServerData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
