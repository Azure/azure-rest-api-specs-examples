import com.azure.core.util.Context;

/** Samples for ResourceGuardProxies Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/recoveryservicesbackup/resource-manager/Microsoft.RecoveryServices/stable/2021-12-01/examples/ResourceGuardProxyCRUD/ListResourceGuardProxy.json
     */
    /**
     * Sample code: Get VaultGuardProxies.
     *
     * @param manager Entry point to RecoveryServicesBackupManager.
     */
    public static void getVaultGuardProxies(
        com.azure.resourcemanager.recoveryservicesbackup.RecoveryServicesBackupManager manager) {
        manager.resourceGuardProxies().get("sampleVault", "SampleResourceGroup", Context.NONE);
    }
}
