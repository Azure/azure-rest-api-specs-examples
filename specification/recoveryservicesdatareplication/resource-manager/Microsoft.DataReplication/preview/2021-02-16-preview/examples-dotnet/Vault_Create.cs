using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.RecoveryServicesDataReplication;
using Azure.ResourceManager.RecoveryServicesDataReplication.Models;
using Azure.ResourceManager.Resources;

// Generated from example definition: specification/recoveryservicesdatareplication/resource-manager/Microsoft.DataReplication/preview/2021-02-16-preview/examples/Vault_Create.json
// this example is just showing the usage of "Vault_Create" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ResourceGroupResource created on azure
// for more information of creating ResourceGroupResource, please refer to the document of ResourceGroupResource
string subscriptionId = "930CEC23-4430-4513-B855-DBA237E2F3BF";
string resourceGroupName = "rgrecoveryservicesdatareplication";
ResourceIdentifier resourceGroupResourceId = ResourceGroupResource.CreateResourceIdentifier(subscriptionId, resourceGroupName);
ResourceGroupResource resourceGroupResource = client.GetResourceGroupResource(resourceGroupResourceId);

// get the collection of this DataReplicationVaultResource
DataReplicationVaultCollection collection = resourceGroupResource.GetDataReplicationVaults();

// invoke the operation
string vaultName = "4";
DataReplicationVaultData data = new DataReplicationVaultData(new AzureLocation("eck"))
{
    Properties = new DataReplicationVaultProperties()
    {
        VaultType = DataReplicationReplicationVaultType.DisasterRecovery,
    },
    Tags =
    {
    ["key5359"] = "ljfilxolxzuxrauopwtyxghrp",
    },
};
ArmOperation<DataReplicationVaultResource> lro = await collection.CreateOrUpdateAsync(WaitUntil.Completed, vaultName, data);
DataReplicationVaultResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
DataReplicationVaultData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
