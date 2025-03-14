using Azure;
using Azure.ResourceManager;
using System;
using System.Collections.Generic;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Models;
using Azure.ResourceManager.Sql.Models;
using Azure.ResourceManager.Sql;

// Generated from example definition: specification/sql/resource-manager/Microsoft.Sql/preview/2023-02-01-preview/examples/PatchVCoreDatabaseWithKeysAndEncryptionProtector.json
// this example is just showing the usage of "Databases_Update" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this SqlDatabaseResource created on azure
// for more information of creating SqlDatabaseResource, please refer to the document of SqlDatabaseResource
string subscriptionId = "00000000-1111-2222-3333-444444444444";
string resourceGroupName = "Default-SQL-SouthEastAsia";
string serverName = "testsvr";
string databaseName = "testdb";
ResourceIdentifier sqlDatabaseResourceId = SqlDatabaseResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, serverName, databaseName);
SqlDatabaseResource sqlDatabase = client.GetSqlDatabaseResource(sqlDatabaseResourceId);

// invoke the operation
SqlDatabasePatch patch = new SqlDatabasePatch
{
    Sku = new SqlSku("S0")
    {
        Tier = "Standard",
    },
    Identity = new DatabaseIdentity
    {
        IdentityType = DatabaseIdentityType.UserAssigned,
        UserAssignedIdentities =
        {
        ["/subscriptions/00000000-1111-2222-3333-444444444444/resourcegroups/Default-SQL-SouthEastAsia/providers/Microsoft.ManagedIdentity/userAssignedIdentities/umi"] = new UserAssignedIdentity(),
        ["/subscriptions/00000000-1111-2222-3333-444444444444/resourcegroups/Default-SQL-SouthEastAsia/providers/Microsoft.ManagedIdentity/userAssignedIdentities/umiToDelete"] = default
        },
    },
    Keys =
    {
    ["https://your-key-vault-name.vault.azure.net/yourKey/yourKeyVersion"] = new SqlDatabaseKey(),
    ["https://your-key-vault-name.vault.azure.net/yourKey2/yourKey2VersionToDelete"] = default
    },
    EncryptionProtector = "https://your-key-vault-name.vault.azure.net/yourKey/yourKeyVersion",
};
ArmOperation<SqlDatabaseResource> lro = await sqlDatabase.UpdateAsync(WaitUntil.Completed, patch);
SqlDatabaseResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
SqlDatabaseData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
