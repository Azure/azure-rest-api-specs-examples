using System;
using System.Collections.Generic;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Sql;
using Azure.ResourceManager.Sql.Models;

// Generated from example definition: specification/sql/resource-manager/Microsoft.Sql/preview/2020-11-01-preview/examples/ManagedDatabaseSensitivityLabelsRecommendedUpdate.json
// this example is just showing the usage of "ManagedDatabaseSensitivityLabels_UpdateRecommended" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ManagedDatabaseResource created on azure
// for more information of creating ManagedDatabaseResource, please refer to the document of ManagedDatabaseResource
string subscriptionId = "00000000-1111-2222-3333-444444444444";
string resourceGroupName = "myRG";
string managedInstanceName = "myManagedInstanceName";
string databaseName = "myDatabase";
ResourceIdentifier managedDatabaseResourceId = ManagedDatabaseResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, managedInstanceName, databaseName);
ManagedDatabaseResource managedDatabase = client.GetManagedDatabaseResource(managedDatabaseResourceId);

// invoke the operation
RecommendedSensitivityLabelUpdateList recommendedSensitivityLabelUpdateList = new RecommendedSensitivityLabelUpdateList()
{
    Operations =
    {
    new RecommendedSensitivityLabelUpdate()
    {
    Op = RecommendedSensitivityLabelUpdateKind.Enable,
    Schema = "dbo",
    Table = "table1",
    Column = "column1",
    },new RecommendedSensitivityLabelUpdate()
    {
    Op = RecommendedSensitivityLabelUpdateKind.Disable,
    Schema = "dbo",
    Table = "table2",
    Column = "column2",
    },new RecommendedSensitivityLabelUpdate()
    {
    Op = RecommendedSensitivityLabelUpdateKind.Disable,
    Schema = "dbo",
    Table = "Table1",
    Column = "Column3",
    }
    },
};
await managedDatabase.UpdateRecommendedManagedDatabaseSensitivityLabelAsync(recommendedSensitivityLabelUpdateList);

Console.WriteLine($"Succeeded");
