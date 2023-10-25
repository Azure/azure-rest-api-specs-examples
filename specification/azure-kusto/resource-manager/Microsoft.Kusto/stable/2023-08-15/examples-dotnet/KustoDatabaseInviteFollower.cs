using System;
using System.Threading.Tasks;
using System.Xml;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Kusto;
using Azure.ResourceManager.Kusto.Models;

// Generated from example definition: specification/azure-kusto/resource-manager/Microsoft.Kusto/stable/2023-08-15/examples/KustoDatabaseInviteFollower.json
// this example is just showing the usage of "Database_InviteFollower" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this KustoDatabaseResource created on azure
// for more information of creating KustoDatabaseResource, please refer to the document of KustoDatabaseResource
string subscriptionId = "12345678-1234-1234-1234-123456789098";
string resourceGroupName = "kustorptest";
string clusterName = "kustoCluster";
string databaseName = "database";
ResourceIdentifier kustoDatabaseResourceId = KustoDatabaseResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, clusterName, databaseName);
KustoDatabaseResource kustoDatabase = client.GetKustoDatabaseResource(kustoDatabaseResourceId);

// invoke the operation
DatabaseInviteFollowerContent content = new DatabaseInviteFollowerContent("invitee@contoso.com")
{
    TableLevelSharingProperties = new KustoDatabaseTableLevelSharingProperties()
    {
        TablesToInclude =
        {
        "Table1"
        },
        TablesToExclude =
        {
        "Table2"
        },
        ExternalTablesToInclude =
        {
        "ExternalTable*"
        },
        ExternalTablesToExclude =
        {
        },
        MaterializedViewsToInclude =
        {
        "MaterializedViewTable1"
        },
        MaterializedViewsToExclude =
        {
        "MaterializedViewTable2"
        },
        FunctionsToInclude =
        {
        "functionsToInclude1"
        },
        FunctionsToExclude =
        {
        "functionsToExclude2"
        },
    },
};
DatabaseInviteFollowerResult result = await kustoDatabase.InviteFollowerDatabaseAsync(content);

Console.WriteLine($"Succeeded: {result}");
