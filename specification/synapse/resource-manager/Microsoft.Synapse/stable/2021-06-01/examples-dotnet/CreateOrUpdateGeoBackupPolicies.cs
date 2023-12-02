using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Synapse;
using Azure.ResourceManager.Synapse.Models;

// Generated from example definition: specification/synapse/resource-manager/Microsoft.Synapse/stable/2021-06-01/examples/CreateOrUpdateGeoBackupPolicies.json
// this example is just showing the usage of "SqlPoolGeoBackupPolicies_CreateOrUpdate" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this SynapseSqlPoolResource created on azure
// for more information of creating SynapseSqlPoolResource, please refer to the document of SynapseSqlPoolResource
string subscriptionId = "00000000-1111-2222-3333-444444444444";
string resourceGroupName = "testrg";
string workspaceName = "testws";
string sqlPoolName = "testdw";
ResourceIdentifier synapseSqlPoolResourceId = SynapseSqlPoolResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, workspaceName, sqlPoolName);
SynapseSqlPoolResource synapseSqlPool = client.GetSynapseSqlPoolResource(synapseSqlPoolResourceId);

// get the collection of this SynapseGeoBackupPolicyResource
SynapseGeoBackupPolicyCollection collection = synapseSqlPool.GetSynapseGeoBackupPolicies();

// invoke the operation
SynapseGeoBackupPolicyName geoBackupPolicyName = SynapseGeoBackupPolicyName.Default;
SynapseGeoBackupPolicyData data = new SynapseGeoBackupPolicyData(SynapseGeoBackupPolicyState.Enabled);
ArmOperation<SynapseGeoBackupPolicyResource> lro = await collection.CreateOrUpdateAsync(WaitUntil.Completed, geoBackupPolicyName, data);
SynapseGeoBackupPolicyResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
SynapseGeoBackupPolicyData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
