using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Synapse;
using Azure.ResourceManager.Synapse.Models;

// Generated from example definition: specification/synapse/resource-manager/Microsoft.Synapse/stable/2021-06-01/examples/DeleteSqlPoolColumnSensitivityLabel.json
// this example is just showing the usage of "SqlPoolSensitivityLabels_Delete" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this SynapseSensitivityLabelResource created on azure
// for more information of creating SynapseSensitivityLabelResource, please refer to the document of SynapseSensitivityLabelResource
string subscriptionId = "00000000-1111-2222-3333-444444444444";
string resourceGroupName = "myRG";
string workspaceName = "myServer";
string sqlPoolName = "myDatabase";
string schemaName = "dbo";
string tableName = "myTable";
string columnName = "myColumn";
ResourceIdentifier synapseSensitivityLabelResourceId = SynapseSensitivityLabelResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, workspaceName, sqlPoolName, schemaName, tableName, columnName);
SynapseSensitivityLabelResource synapseSensitivityLabel = client.GetSynapseSensitivityLabelResource(synapseSensitivityLabelResourceId);

// invoke the operation
await synapseSensitivityLabel.DeleteAsync(WaitUntil.Completed);

Console.WriteLine($"Succeeded");
