using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Synapse;
using Azure.ResourceManager.Synapse.Models;

// Generated from example definition: specification/synapse/resource-manager/Microsoft.Synapse/stable/2021-06-01/examples/SensitivityLabelsRecommendedUpdate.json
// this example is just showing the usage of "SqlPoolRecommendedSensitivityLabels_Update" operation, for the dependent resources, they will have to be created separately.

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
SynapseRecommendedSensitivityLabelUpdateOperationListResult synapseRecommendedSensitivityLabelUpdateOperationListResult = new SynapseRecommendedSensitivityLabelUpdateOperationListResult()
{
    Operations =
    {
    new SynapseRecommendedSensitivityLabelUpdate()
    {
    Op = SynapseRecommendedSensitivityLabelUpdateKind.Enable,
    Schema = "dbo",
    Table = "table1",
    Column = "column1",
    },new SynapseRecommendedSensitivityLabelUpdate()
    {
    Op = SynapseRecommendedSensitivityLabelUpdateKind.Enable,
    Schema = "dbo",
    Table = "table2",
    Column = "column2",
    },new SynapseRecommendedSensitivityLabelUpdate()
    {
    Op = SynapseRecommendedSensitivityLabelUpdateKind.Disable,
    Schema = "dbo",
    Table = "table1",
    Column = "column3",
    }
    },
};
await synapseSqlPool.UpdateSqlPoolRecommendedSensitivityLabelAsync(synapseRecommendedSensitivityLabelUpdateOperationListResult);

Console.WriteLine($"Succeeded");
