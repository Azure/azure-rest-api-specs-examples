using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Synapse;

// Generated from example definition: specification/synapse/resource-manager/Microsoft.Synapse/stable/2021-06-01/examples/DeleteSqlPoolWorkloadGroupWorkloadClassifer.json
// this example is just showing the usage of "SqlPoolWorkloadClassifier_Delete" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this SynapseWorkloadClassifierResource created on azure
// for more information of creating SynapseWorkloadClassifierResource, please refer to the document of SynapseWorkloadClassifierResource
string subscriptionId = "00000000-1111-2222-3333-444444444444";
string resourceGroupName = "sqlcrudtest-6852";
string workspaceName = "sqlcrudtest-2080";
string sqlPoolName = "sqlcrudtest-9187";
string workloadGroupName = "wlm_workloadgroup";
string workloadClassifierName = "wlm_workloadclassifier";
ResourceIdentifier synapseWorkloadClassifierResourceId = SynapseWorkloadClassifierResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, workspaceName, sqlPoolName, workloadGroupName, workloadClassifierName);
SynapseWorkloadClassifierResource synapseWorkloadClassifier = client.GetSynapseWorkloadClassifierResource(synapseWorkloadClassifierResourceId);

// invoke the operation
await synapseWorkloadClassifier.DeleteAsync(WaitUntil.Completed);

Console.WriteLine("Succeeded");
