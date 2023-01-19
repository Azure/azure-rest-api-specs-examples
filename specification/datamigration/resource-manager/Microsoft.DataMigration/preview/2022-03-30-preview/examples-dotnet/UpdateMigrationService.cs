using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.DataMigration;
using Azure.ResourceManager.DataMigration.Models;
using Azure.ResourceManager.Resources;

// Generated from example definition: specification/datamigration/resource-manager/Microsoft.DataMigration/preview/2022-03-30-preview/examples/UpdateMigrationService.json
// this example is just showing the usage of "SqlMigrationServices_Update" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this SqlMigrationServiceResource created on azure
// for more information of creating SqlMigrationServiceResource, please refer to the document of SqlMigrationServiceResource
string subscriptionId = "00000000-1111-2222-3333-444444444444";
string resourceGroupName = "testrg";
string sqlMigrationServiceName = "testagent";
ResourceIdentifier sqlMigrationServiceResourceId = SqlMigrationServiceResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, sqlMigrationServiceName);
SqlMigrationServiceResource sqlMigrationService = client.GetSqlMigrationServiceResource(sqlMigrationServiceResourceId);

// invoke the operation
SqlMigrationServicePatch patch = new SqlMigrationServicePatch()
{
    Tags =
    {
    ["mytag"] = "myval",
    },
};
ArmOperation<SqlMigrationServiceResource> lro = await sqlMigrationService.UpdateAsync(WaitUntil.Completed, patch);
SqlMigrationServiceResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
SqlMigrationServiceData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
