using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Synapse;
using Azure.ResourceManager.Synapse.Models;

// Generated from example definition: specification/synapse/resource-manager/Microsoft.Synapse/stable/2021-06-01/examples/SensitivityLabelsCurrentUpdate.json
// this example is just showing the usage of "SqlPoolSensitivityLabels_Update" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this SynapseSqlPoolResource created on azure
// for more information of creating SynapseSqlPoolResource, please refer to the document of SynapseSqlPoolResource
string subscriptionId = "00000000-1111-2222-3333-444444444444";
string resourceGroupName = "myRG";
string workspaceName = "myWorkspace";
string sqlPoolName = "mySqlPool";
ResourceIdentifier synapseSqlPoolResourceId = SynapseSqlPoolResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, workspaceName, sqlPoolName);
SynapseSqlPoolResource synapseSqlPool = client.GetSynapseSqlPoolResource(synapseSqlPoolResourceId);

// invoke the operation
SynapseSensitivityLabelUpdateListResult synapseSensitivityLabelUpdateListResult = new SynapseSensitivityLabelUpdateListResult()
{
    Operations =
    {
    new SynapseSensitivityLabelUpdate()
    {
    Op = SynapseSensitivityLabelUpdateKind.Set,
    Schema = "dbo",
    Table = "table1",
    Column = "column1",
    SensitivityLabel = new SynapseSensitivityLabelData()
    {
    LabelName = "Highly Confidential",
    LabelId = Guid.Parse("3A477B16-9423-432B-AA97-6069B481CEC3"),
    InformationType = "Financial",
    InformationTypeId = Guid.Parse("1D3652D6-422C-4115-82F1-65DAEBC665C8"),
    Rank = SynapseSensitivityLabelRank.Low,
    },
    },new SynapseSensitivityLabelUpdate()
    {
    Op = SynapseSensitivityLabelUpdateKind.Set,
    Schema = "dbo",
    Table = "table2",
    Column = "column2",
    SensitivityLabel = new SynapseSensitivityLabelData()
    {
    LabelName = "PII",
    LabelId = Guid.Parse("bf91e08c-f4f0-478a-b016-25164b2a65ff"),
    InformationType = "PhoneNumber",
    InformationTypeId = Guid.Parse("d22fa6e9-5ee4-3bde-4c2b-a409604c4646"),
    Rank = SynapseSensitivityLabelRank.Critical,
    },
    },new SynapseSensitivityLabelUpdate()
    {
    Op = SynapseSensitivityLabelUpdateKind.Remove,
    Schema = "dbo",
    Table = "Table1",
    Column = "Column3",
    }
    },
};
await synapseSqlPool.UpdateSqlPoolSensitivityLabelAsync(synapseSensitivityLabelUpdateListResult);

Console.WriteLine($"Succeeded");
