using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.RecoveryServicesDataReplication.Models;
using Azure.ResourceManager.RecoveryServicesDataReplication;

// Generated from example definition: specification/recoveryservicesdatareplication/resource-manager/Microsoft.DataReplication/preview/2021-02-16-preview/examples/Policy_Get.json
// this example is just showing the usage of "Policy_Get" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this DataReplicationVaultResource created on azure
// for more information of creating DataReplicationVaultResource, please refer to the document of DataReplicationVaultResource
string subscriptionId = "930CEC23-4430-4513-B855-DBA237E2F3BF";
string resourceGroupName = "rgrecoveryservicesdatareplication";
string vaultName = "4";
ResourceIdentifier dataReplicationVaultResourceId = DataReplicationVaultResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, vaultName);
DataReplicationVaultResource dataReplicationVault = client.GetDataReplicationVaultResource(dataReplicationVaultResourceId);

// get the collection of this DataReplicationPolicyResource
DataReplicationPolicyCollection collection = dataReplicationVault.GetDataReplicationPolicies();

// invoke the operation
string policyName = "wdqsacasc";
NullableResponse<DataReplicationPolicyResource> response = await collection.GetIfExistsAsync(policyName);
DataReplicationPolicyResource result = response.HasValue ? response.Value : null;

if (result == null)
{
    Console.WriteLine("Succeeded with null as result");
}
else
{
    // the variable result is a resource, you could call other operations on this instance as well
    // but just for demo, we get its data from this resource instance
    DataReplicationPolicyData resourceData = result.Data;
    // for demo we just print out the id
    Console.WriteLine($"Succeeded on id: {resourceData.Id}");
}
