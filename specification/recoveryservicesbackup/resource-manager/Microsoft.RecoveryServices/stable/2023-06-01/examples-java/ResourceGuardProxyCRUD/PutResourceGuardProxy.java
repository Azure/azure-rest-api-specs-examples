
import com.azure.resourcemanager.recoveryservicesbackup.models.ResourceGuardProxyBase;

/**
 * Samples for ResourceGuardProxyOperation Put.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/recoveryservicesbackup/resource-manager/Microsoft.RecoveryServices/stable/2023-06-01/examples/
     * ResourceGuardProxyCRUD/PutResourceGuardProxy.json
     */
    /**
     * Sample code: Create ResourceGuardProxy.
     * 
     * @param manager Entry point to RecoveryServicesBackupManager.
     */
    public static void createResourceGuardProxy(
        com.azure.resourcemanager.recoveryservicesbackup.RecoveryServicesBackupManager manager) {
        manager.resourceGuardProxyOperations().define("swaggerExample").withRegion((String) null)
            .withExistingVault("sampleVault", "SampleResourceGroup")
            .withProperties(new ResourceGuardProxyBase().withResourceGuardResourceId(
                "/subscriptions/c999d45b-944f-418c-a0d8-c3fcfd1802c8/resourceGroups/vaultguardRGNew/providers/Microsoft.DataProtection/resourceGuards/VaultGuardTestNew"))
            .create();
    }
}
