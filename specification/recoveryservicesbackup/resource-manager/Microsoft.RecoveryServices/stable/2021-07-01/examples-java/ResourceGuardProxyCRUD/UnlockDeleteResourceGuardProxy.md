Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-recoveryservicesbackup_1.0.0-beta.2/sdk/recoveryservicesbackup/azure-resourcemanager-recoveryservicesbackup/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;
import com.azure.resourcemanager.recoveryservicesbackup.models.UnlockDeleteRequest;
import java.util.Arrays;

/** Samples for ResourceGuardProxyOperation UnlockDelete. */
public final class Main {
    /*
     * x-ms-original-file: specification/recoveryservicesbackup/resource-manager/Microsoft.RecoveryServices/stable/2021-07-01/examples/ResourceGuardProxyCRUD/UnlockDeleteResourceGuardProxy.json
     */
    /**
     * Sample code: UnlockDelete ResourceGuardProxy.
     *
     * @param manager Entry point to RecoveryServicesBackupManager.
     */
    public static void unlockDeleteResourceGuardProxy(
        com.azure.resourcemanager.recoveryservicesbackup.RecoveryServicesBackupManager manager) {
        manager
            .resourceGuardProxyOperations()
            .unlockDeleteWithResponse(
                "sampleVault",
                "SampleResourceGroup",
                "swaggerExample",
                new UnlockDeleteRequest()
                    .withResourceGuardOperationRequests(
                        Arrays
                            .asList(
                                "/subscriptions/c999d45b-944f-418c-a0d8-c3fcfd1802c8/resourceGroups/vaultguardRGNew/providers/Microsoft.DataProtection/resourceGuards/VaultGuardTestNew/deleteProtectedItemRequests/default"))
                    .withResourceToBeDeleted(
                        "/subscriptions/62b829ee-7936-40c9-a1c9-47a93f9f3965/resourceGroups/gaallarg/providers/Microsoft.RecoveryServices/vaults/MercuryCrrVault/backupFabrics/Azure/protectionContainers/VMAppContainer;compute;crrtestrg;crrtestvm/protectedItems/SQLDataBase;mssqlserver;testdb"),
                Context.NONE);
    }
}
```
