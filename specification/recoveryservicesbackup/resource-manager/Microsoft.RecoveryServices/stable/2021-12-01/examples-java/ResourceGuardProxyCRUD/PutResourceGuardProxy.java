import com.azure.core.util.Context;

/** Samples for ResourceGuardProxyOperation Put. */
public final class Main {
    /*
     * x-ms-original-file: specification/recoveryservicesbackup/resource-manager/Microsoft.RecoveryServices/stable/2021-12-01/examples/ResourceGuardProxyCRUD/PutResourceGuardProxy.json
     */
    /**
     * Sample code: Create ResourceGuardProxy.
     *
     * @param manager Entry point to RecoveryServicesBackupManager.
     */
    public static void createResourceGuardProxy(
        com.azure.resourcemanager.recoveryservicesbackup.RecoveryServicesBackupManager manager) {
        manager
            .resourceGuardProxyOperations()
            .putWithResponse("sampleVault", "SampleResourceGroup", "swaggerExample", Context.NONE);
    }
}
