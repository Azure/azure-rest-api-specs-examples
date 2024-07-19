using Azure;
using Azure.ResourceManager;
using System;
using System.Collections.Generic;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Sql.Models;
using Azure.ResourceManager.Sql;

// Generated from example definition: specification/sql/resource-manager/Microsoft.Sql/preview/2020-11-01-preview/examples/ManagedDatabaseSecurityEventsGetMax.json
// this example is just showing the usage of "ManagedDatabaseSecurityEvents_ListByDatabase" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ManagedDatabaseResource created on azure
// for more information of creating ManagedDatabaseResource, please refer to the document of ManagedDatabaseResource
string subscriptionId = "00000000-1111-2222-3333-444444444444";
string resourceGroupName = "testrg";
string managedInstanceName = "testcl";
string databaseName = "database1";
ResourceIdentifier managedDatabaseResourceId = ManagedDatabaseResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, managedInstanceName, databaseName);
ManagedDatabaseResource managedDatabase = client.GetManagedDatabaseResource(managedDatabaseResourceId);

// invoke the operation and iterate over the result
string filter = "ShowServerRecords eq true";
int? skip = 0;
int? top = 1;
string skiptoken = "eyJCbG9iTmFtZURhdGVUaW1lIjoiXC9EYXRlKDE1MTIyODg4MTIwMTArMDIwMClcLyIsIkJsb2JOYW1lUm9sbG92ZXJJbmRleCI6IjAiLCJFbmREYXRlIjoiXC9EYXRlKDE1MTI0NjYyMDA1MjkpXC8iLCJJc1NraXBUb2tlblNldCI6ZmFsc2UsIklzVjJCbG9iVGltZUZvcm1hdCI6dHJ1ZSwiU2hvd1NlcnZlclJlY29yZHMiOmZhbHNlLCJTa2lwVmFsdWUiOjAsIlRha2VWYWx1ZSI6MTB9";
await foreach (SecurityEvent item in managedDatabase.GetManagedDatabaseSecurityEventsByDatabaseAsync(filter: filter, skip: skip, top: top, skiptoken: skiptoken))
{
    Console.WriteLine($"Succeeded: {item}");
}

Console.WriteLine($"Succeeded");
