using System;
using System.Collections.Generic;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Sql;
using Azure.ResourceManager.Sql.Models;

// Generated from example definition: specification/sql/resource-manager/Microsoft.Sql/preview/2020-11-01-preview/examples/ManagedDatabaseSensitivityLabelsCurrentUpdate.json
// this example is just showing the usage of "ManagedDatabaseSensitivityLabels_Update" operation, for the dependent resources, they will have to be created separately.

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
SensitivityLabelUpdateList sensitivityLabelUpdateList = new SensitivityLabelUpdateList()
{
    Operations =
    {
    new SensitivityLabelUpdate()
    {
    Op = SensitivityLabelUpdateKind.Set,
    Schema = "dbo",
    Table = "table1",
    Column = "column1",
    SensitivityLabel = new SensitivityLabelData()
    {
    LabelName = "Highly Confidential",
    LabelId = "3A477B16-9423-432B-AA97-6069B481CEC3",
    InformationType = "Financial",
    InformationTypeId = "1D3652D6-422C-4115-82F1-65DAEBC665C8",
    },
    },new SensitivityLabelUpdate()
    {
    Op = SensitivityLabelUpdateKind.Set,
    Schema = "dbo",
    Table = "table2",
    Column = "column2",
    SensitivityLabel = new SensitivityLabelData()
    {
    LabelName = "PII",
    LabelId = "bf91e08c-f4f0-478a-b016-25164b2a65ff",
    InformationType = "PhoneNumber",
    InformationTypeId = "d22fa6e9-5ee4-3bde-4c2b-a409604c4646",
    },
    },new SensitivityLabelUpdate()
    {
    Op = SensitivityLabelUpdateKind.Remove,
    Schema = "dbo",
    Table = "Table1",
    Column = "Column3",
    }
    },
};
await managedDatabase.UpdateManagedDatabaseSensitivityLabelAsync(sensitivityLabelUpdateList);

Console.WriteLine($"Succeeded");
