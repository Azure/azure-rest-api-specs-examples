using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Synapse;
using Azure.ResourceManager.Synapse.Models;

// Generated from example definition: specification/synapse/resource-manager/Microsoft.Synapse/stable/2021-06-01/examples/WorkspaceManagedSqlServerEncryptionProtectorCreateOrUpdateKeyVault.json
// this example is just showing the usage of "WorkspaceManagedSqlServerEncryptionProtector_CreateOrUpdate" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this SynapseEncryptionProtectorResource created on azure
// for more information of creating SynapseEncryptionProtectorResource, please refer to the document of SynapseEncryptionProtectorResource
string subscriptionId = "00000000-1111-2222-3333-444444444444";
string resourceGroupName = "wsg-7398";
string workspaceName = "testWorkspace";
SynapseEncryptionProtectorName encryptionProtectorName = SynapseEncryptionProtectorName.Current;
ResourceIdentifier synapseEncryptionProtectorResourceId = SynapseEncryptionProtectorResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, workspaceName, encryptionProtectorName);
SynapseEncryptionProtectorResource synapseEncryptionProtector = client.GetSynapseEncryptionProtectorResource(synapseEncryptionProtectorResourceId);

// invoke the operation
SynapseEncryptionProtectorData data = new SynapseEncryptionProtectorData()
{
    ServerKeyName = "someVault_someKey_01234567890123456789012345678901",
    ServerKeyType = SynapseServerKeyType.AzureKeyVault,
};
ArmOperation<SynapseEncryptionProtectorResource> lro = await synapseEncryptionProtector.UpdateAsync(WaitUntil.Completed, data);
SynapseEncryptionProtectorResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
SynapseEncryptionProtectorData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
