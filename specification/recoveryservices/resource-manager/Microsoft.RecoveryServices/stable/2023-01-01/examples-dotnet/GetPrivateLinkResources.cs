using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.RecoveryServices;

// Generated from example definition: specification/recoveryservices/resource-manager/Microsoft.RecoveryServices/stable/2023-01-01/examples/GetPrivateLinkResources.json
// this example is just showing the usage of "PrivateLinkResources_Get" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this RecoveryServicesVaultResource created on azure
// for more information of creating RecoveryServicesVaultResource, please refer to the document of RecoveryServicesVaultResource
string subscriptionId = "6c48fa17-39c7-45f1-90ac-47a587128ace";
string resourceGroupName = "petesting";
string vaultName = "pemsi-ecy-rsv2";
ResourceIdentifier recoveryServicesVaultResourceId = RecoveryServicesVaultResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, vaultName);
RecoveryServicesVaultResource recoveryServicesVault = client.GetRecoveryServicesVaultResource(recoveryServicesVaultResourceId);

// get the collection of this RecoveryServicesPrivateLinkResource
RecoveryServicesPrivateLinkResourceCollection collection = recoveryServicesVault.GetRecoveryServicesPrivateLinkResources();

// invoke the operation
string privateLinkResourceName = "backupResource";
bool result = await collection.ExistsAsync(privateLinkResourceName);

Console.WriteLine($"Succeeded: {result}");
