using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.DataMigration;
using Azure.ResourceManager.DataMigration.Models;

// Generated from example definition: specification/datamigration/resource-manager/Microsoft.DataMigration/preview/2022-03-30-preview/examples/SqlVmGetDatabaseMigrationExpanded.json
// this example is just showing the usage of "DatabaseMigrationsSqlVm_Get" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this DatabaseMigrationSqlVmResource created on azure
// for more information of creating DatabaseMigrationSqlVmResource, please refer to the document of DatabaseMigrationSqlVmResource
string subscriptionId = "00000000-1111-2222-3333-444444444444";
string resourceGroupName = "testrg";
string sqlVirtualMachineName = "testvm";
string targetDBName = "db1";
ResourceIdentifier databaseMigrationSqlVmResourceId = DatabaseMigrationSqlVmResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, sqlVirtualMachineName, targetDBName);
DatabaseMigrationSqlVmResource databaseMigrationSqlVm = client.GetDatabaseMigrationSqlVmResource(databaseMigrationSqlVmResourceId);

// invoke the operation
string expand = "MigrationStatusDetails";
DatabaseMigrationSqlVmResource result = await databaseMigrationSqlVm.GetAsync(expand: expand);

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
DatabaseMigrationSqlVmData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
