using Azure;
using Azure.ResourceManager;
using System;
using System.Collections.Generic;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Sql.Models;
using Azure.ResourceManager.Sql;

// Generated from example definition: specification/sql/resource-manager/Microsoft.Sql/preview/2024-11-01-preview/examples/ManagedInstanceQueryStatisticsList.json
// this example is just showing the usage of "ManagedDatabaseQueries_ListByQuery" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ManagedDatabaseResource created on azure
// for more information of creating ManagedDatabaseResource, please refer to the document of ManagedDatabaseResource
string subscriptionId = "00000000-1111-2222-3333-444444444444";
string resourceGroupName = "sqlcrudtest-7398";
string managedInstanceName = "sqlcrudtest-4645";
string databaseName = "database_1";
ResourceIdentifier managedDatabaseResourceId = ManagedDatabaseResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, managedInstanceName, databaseName);
ManagedDatabaseResource managedDatabase = client.GetManagedDatabaseResource(managedDatabaseResourceId);

// invoke the operation and iterate over the result
string queryId = "42";
await foreach (QueryStatistics item in managedDatabase.GetQueryStatisticsAsync(queryId))
{
    Console.WriteLine($"Succeeded: {item}");
}

Console.WriteLine("Succeeded");
