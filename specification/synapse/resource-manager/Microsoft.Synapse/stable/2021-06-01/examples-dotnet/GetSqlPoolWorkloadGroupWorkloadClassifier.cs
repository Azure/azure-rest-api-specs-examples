using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Synapse;

// Generated from example definition: specification/synapse/resource-manager/Microsoft.Synapse/stable/2021-06-01/examples/GetSqlPoolWorkloadGroupWorkloadClassifier.json
// this example is just showing the usage of "SqlPoolWorkloadClassifier_Get" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this SynapseWorkloadGroupResource created on azure
// for more information of creating SynapseWorkloadGroupResource, please refer to the document of SynapseWorkloadGroupResource
string subscriptionId = "00000000-1111-2222-3333-444444444444";
string resourceGroupName = "sqlcrudtest-6852";
string workspaceName = "sqlcrudtest-2080";
string sqlPoolName = "sqlcrudtest-9187";
string workloadGroupName = "wlm_workloadgroup";
ResourceIdentifier synapseWorkloadGroupResourceId = SynapseWorkloadGroupResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, workspaceName, sqlPoolName, workloadGroupName);
SynapseWorkloadGroupResource synapseWorkloadGroup = client.GetSynapseWorkloadGroupResource(synapseWorkloadGroupResourceId);

// get the collection of this SynapseWorkloadClassifierResource
SynapseWorkloadClassifierCollection collection = synapseWorkloadGroup.GetSynapseWorkloadClassifiers();

// invoke the operation
string workloadClassifierName = "wlm_classifier";
bool result = await collection.ExistsAsync(workloadClassifierName);

Console.WriteLine($"Succeeded: {result}");
