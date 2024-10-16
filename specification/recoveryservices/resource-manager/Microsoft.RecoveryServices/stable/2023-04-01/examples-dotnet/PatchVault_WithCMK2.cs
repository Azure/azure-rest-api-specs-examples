using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Models;
using Azure.ResourceManager.RecoveryServices.Models;
using Azure.ResourceManager.Resources;
using Azure.ResourceManager.RecoveryServices;

// Generated from example definition: specification/recoveryservices/resource-manager/Microsoft.RecoveryServices/stable/2023-04-01/examples/PatchVault_WithCMK2.json
// this example is just showing the usage of "Vaults_Update" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this RecoveryServicesVaultResource created on azure
// for more information of creating RecoveryServicesVaultResource, please refer to the document of RecoveryServicesVaultResource
string subscriptionId = "77777777-b0c6-47a2-b37c-d8e65a629c18";
string resourceGroupName = "HelloWorld";
string vaultName = "swaggerExample";
ResourceIdentifier recoveryServicesVaultResourceId = RecoveryServicesVaultResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, vaultName);
RecoveryServicesVaultResource recoveryServicesVault = client.GetRecoveryServicesVaultResource(recoveryServicesVaultResourceId);

// invoke the operation
RecoveryServicesVaultPatch patch = new RecoveryServicesVaultPatch(new AzureLocation("placeholder"))
{
    Properties = new RecoveryServicesVaultProperties()
    {
        Encryption = new VaultPropertiesEncryption()
        {
            KekIdentity = new CmkKekIdentity()
            {
                UseSystemAssignedIdentity = true,
            },
        },
    },
    Identity = new ManagedServiceIdentity("SystemAssigned"),
    Tags =
    {
    ["PatchKey"] = "PatchKeyUpdated",
    },
};
ArmOperation<RecoveryServicesVaultResource> lro = await recoveryServicesVault.UpdateAsync(WaitUntil.Completed, patch);
RecoveryServicesVaultResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
RecoveryServicesVaultData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
