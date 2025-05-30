using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Synapse.Models;
using Azure.ResourceManager.Synapse;

// Generated from example definition: specification/synapse/resource-manager/Microsoft.Synapse/stable/2021-06-01/examples/ExtendedSqlPoolBlobAuditingCreateMax.json
// this example is just showing the usage of "ExtendedSqlPoolBlobAuditingPolicies_CreateOrUpdate" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this SynapseExtendedSqlPoolBlobAuditingPolicyResource created on azure
// for more information of creating SynapseExtendedSqlPoolBlobAuditingPolicyResource, please refer to the document of SynapseExtendedSqlPoolBlobAuditingPolicyResource
string subscriptionId = "00000000-1111-2222-3333-444444444444";
string resourceGroupName = "blobauditingtest-4799";
string workspaceName = "blobauditingtest-6440";
string sqlPoolName = "testdb";
ResourceIdentifier synapseExtendedSqlPoolBlobAuditingPolicyResourceId = SynapseExtendedSqlPoolBlobAuditingPolicyResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, workspaceName, sqlPoolName);
SynapseExtendedSqlPoolBlobAuditingPolicyResource synapseExtendedSqlPoolBlobAuditingPolicy = client.GetSynapseExtendedSqlPoolBlobAuditingPolicyResource(synapseExtendedSqlPoolBlobAuditingPolicyResourceId);

// invoke the operation
SynapseExtendedSqlPoolBlobAuditingPolicyData data = new SynapseExtendedSqlPoolBlobAuditingPolicyData
{
    PredicateExpression = "statement = 'select 1'",
    State = SynapseBlobAuditingPolicyState.Enabled,
    StorageEndpoint = "https://mystorage.blob.core.windows.net",
    StorageAccountAccessKey = "sdlfkjabc+sdlfkjsdlkfsjdfLDKFTERLKFDFKLjsdfksjdflsdkfD2342309432849328476458/3RSD==",
    RetentionDays = 6,
    AuditActionsAndGroups = { "DATABASE_LOGOUT_GROUP", "DATABASE_ROLE_MEMBER_CHANGE_GROUP", "UPDATE on database::TestDatabaseName by public" },
    StorageAccountSubscriptionId = Guid.Parse("00000000-1234-0000-5678-000000000000"),
    IsStorageSecondaryKeyInUse = false,
    IsAzureMonitorTargetEnabled = true,
};
ArmOperation<SynapseExtendedSqlPoolBlobAuditingPolicyResource> lro = await synapseExtendedSqlPoolBlobAuditingPolicy.CreateOrUpdateAsync(WaitUntil.Completed, data);
SynapseExtendedSqlPoolBlobAuditingPolicyResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
SynapseExtendedSqlPoolBlobAuditingPolicyData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
