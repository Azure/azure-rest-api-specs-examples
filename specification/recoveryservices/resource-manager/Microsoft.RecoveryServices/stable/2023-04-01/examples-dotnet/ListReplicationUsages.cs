using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Models;
using Azure.ResourceManager.RecoveryServices;
using Azure.ResourceManager.RecoveryServices.Models;
using Azure.ResourceManager.Resources;

// Generated from example definition: specification/recoveryservices/resource-manager/Microsoft.RecoveryServices/stable/2023-04-01/examples/ListReplicationUsages.json
// this example is just showing the usage of "ReplicationUsages_List" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this RecoveryServicesVaultResource created on azure
// for more information of creating RecoveryServicesVaultResource, please refer to the document of RecoveryServicesVaultResource
string subscriptionId = "6808dbbc-98c7-431f-a1b1-9580902423b7";
string resourceGroupName = "avrai7517RG1";
string vaultName = "avrai7517Vault1";
ResourceIdentifier recoveryServicesVaultResourceId = RecoveryServicesVaultResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, vaultName);
RecoveryServicesVaultResource recoveryServicesVault = client.GetRecoveryServicesVaultResource(recoveryServicesVaultResourceId);

// invoke the operation and iterate over the result
await foreach (ReplicationUsage item in recoveryServicesVault.GetReplicationUsagesAsync())
{
    Console.WriteLine($"Succeeded: {item}");
}

Console.WriteLine($"Succeeded");
