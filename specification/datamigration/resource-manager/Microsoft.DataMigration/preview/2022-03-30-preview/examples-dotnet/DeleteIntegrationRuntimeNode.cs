using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.DataMigration;
using Azure.ResourceManager.DataMigration.Models;
using Azure.ResourceManager.Resources;

// Generated from example definition: specification/datamigration/resource-manager/Microsoft.DataMigration/preview/2022-03-30-preview/examples/DeleteIntegrationRuntimeNode.json
// this example is just showing the usage of "SqlMigrationServices_deleteNode" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this SqlMigrationServiceResource created on azure
// for more information of creating SqlMigrationServiceResource, please refer to the document of SqlMigrationServiceResource
string subscriptionId = "00000000-1111-2222-3333-444444444444";
string resourceGroupName = "testrg";
string sqlMigrationServiceName = "service1";
ResourceIdentifier sqlMigrationServiceResourceId = SqlMigrationServiceResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, sqlMigrationServiceName);
SqlMigrationServiceResource sqlMigrationService = client.GetSqlMigrationServiceResource(sqlMigrationServiceResourceId);

// invoke the operation
DeleteNode deleteNode = new DeleteNode()
{
    NodeName = "nodeName",
    IntegrationRuntimeName = "IRName",
};
DeleteNode result = await sqlMigrationService.DeleteNodeAsync(deleteNode);

Console.WriteLine($"Succeeded: {result}");
