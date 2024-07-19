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

// Generated from example definition: specification/sql/resource-manager/Microsoft.Sql/preview/2021-02-01-preview/examples/ListDatabaseExtensions.json
// this example is just showing the usage of "DatabaseExtensions_ListByDatabase" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this SqlDatabaseResource created on azure
// for more information of creating SqlDatabaseResource, please refer to the document of SqlDatabaseResource
string subscriptionId = "7b2515fe-f230-4017-8cf0-695163acab85";
string resourceGroupName = "rg_4007c5a9-b3b0-41e1-bd46-9eef38768a4a";
string serverName = "srv_3b67ec2a-519b-43a7-8533-fb62dce3431e";
string databaseName = "719d8fa4-bf0f-48fc-8cd3-ef40fe6ba1fe";
ResourceIdentifier sqlDatabaseResourceId = SqlDatabaseResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, serverName, databaseName);
SqlDatabaseResource sqlDatabase = client.GetSqlDatabaseResource(sqlDatabaseResourceId);

// invoke the operation and iterate over the result
await foreach (ImportExportExtensionsOperationResult item in sqlDatabase.GetDatabaseExtensionsAsync())
{
    Console.WriteLine($"Succeeded: {item}");
}

Console.WriteLine($"Succeeded");
