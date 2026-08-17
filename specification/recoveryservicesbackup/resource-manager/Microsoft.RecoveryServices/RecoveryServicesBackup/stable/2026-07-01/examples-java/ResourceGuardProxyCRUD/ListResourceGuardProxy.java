
/**
 * Samples for ResourceGuardProxyOperation List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-07-01/ResourceGuardProxyCRUD/ListResourceGuardProxy.json
     */
    /**
     * Sample code: Get VaultGuardProxies.
     * 
     * @param manager Entry point to RecoveryServicesBackupManager.
     */
    public static void
        getVaultGuardProxies(com.azure.resourcemanager.recoveryservicesbackup.RecoveryServicesBackupManager manager) {
        manager.resourceGuardProxyOperations().list("sampleVault", "SampleResourceGroup",
            com.azure.core.util.Context.NONE);
    }
}
