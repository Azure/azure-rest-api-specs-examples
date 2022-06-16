import com.azure.core.util.Context;

/** Samples for ResourceGuardProxyOperation Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/recoveryservicesbackup/resource-manager/Microsoft.RecoveryServices/stable/2021-12-01/examples/ResourceGuardProxyCRUD/DeleteResourceGuardProxy.json
     */
    /**
     * Sample code: Delete ResourceGuardProxy.
     *
     * @param manager Entry point to RecoveryServicesBackupManager.
     */
    public static void deleteResourceGuardProxy(
        com.azure.resourcemanager.recoveryservicesbackup.RecoveryServicesBackupManager manager) {
        manager
            .resourceGuardProxyOperations()
            .deleteWithResponse("sampleVault", "SampleResourceGroup", "swaggerExample", Context.NONE);
    }
}
