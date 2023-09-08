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

// Generated from example definition: specification/recoveryservices/resource-manager/Microsoft.RecoveryServices/stable/2023-04-01/examples/PATCHVault_WithUserAssignedIdentity.json
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
    Identity = new ManagedServiceIdentity("UserAssigned")
    {
        UserAssignedIdentities =
        {
        [new ResourceIdentifier("/subscriptions/85bf5e8c-3084-4f42-add2-746ebb7e97b2/resourcegroups/defaultrg/providers/Microsoft.ManagedIdentity/userAssignedIdentities/examplemsi")] = new UserAssignedIdentity(),
        },
    },
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
