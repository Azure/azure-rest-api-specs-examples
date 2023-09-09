using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.RecoveryServicesBackup;
using Azure.ResourceManager.RecoveryServicesBackup.Models;

// Generated from example definition: specification/recoveryservicesbackup/resource-manager/Microsoft.RecoveryServices/stable/2023-04-01/examples/ResourceGuardProxyCRUD/UnlockDeleteResourceGuardProxy.json
// this example is just showing the usage of "ResourceGuardProxy_UnlockDelete" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ResourceGuardProxyResource created on azure
// for more information of creating ResourceGuardProxyResource, please refer to the document of ResourceGuardProxyResource
string subscriptionId = "0b352192-dcac-4cc7-992e-a96190ccc68c";
string resourceGroupName = "SampleResourceGroup";
string vaultName = "sampleVault";
string resourceGuardProxyName = "swaggerExample";
ResourceIdentifier resourceGuardProxyResourceId = ResourceGuardProxyResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, vaultName, resourceGuardProxyName);
ResourceGuardProxyResource resourceGuardProxy = client.GetResourceGuardProxyResource(resourceGuardProxyResourceId);

// invoke the operation
UnlockDeleteContent content = new UnlockDeleteContent()
{
    ResourceGuardOperationRequests =
    {
    "/subscriptions/c999d45b-944f-418c-a0d8-c3fcfd1802c8/resourceGroups/vaultguardRGNew/providers/Microsoft.DataProtection/resourceGuards/VaultGuardTestNew/deleteProtectedItemRequests/default"
    },
    ResourceToBeDeleted = "/subscriptions/62b829ee-7936-40c9-a1c9-47a93f9f3965/resourceGroups/gaallarg/providers/Microsoft.RecoveryServices/vaults/MercuryCrrVault/backupFabrics/Azure/protectionContainers/VMAppContainer;compute;crrtestrg;crrtestvm/protectedItems/SQLDataBase;mssqlserver;testdb",
};
UnlockDeleteResult result = await resourceGuardProxy.UnlockDeleteAsync(content);

Console.WriteLine($"Succeeded: {result}");
