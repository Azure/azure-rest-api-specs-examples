/** Samples for ResourceGuardProxyOperation Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/recoveryservicesbackup/resource-manager/Microsoft.RecoveryServices/stable/2023-04-01/examples/ResourceGuardProxyCRUD/GetResourceGuardProxy.json
     */
    /**
     * Sample code: Get ResourceGuardProxy.
     *
     * @param manager Entry point to RecoveryServicesBackupManager.
     */
    public static void getResourceGuardProxy(
        com.azure.resourcemanager.recoveryservicesbackup.RecoveryServicesBackupManager manager) {
        manager
            .resourceGuardProxyOperations()
            .getWithResponse("sampleVault", "SampleResourceGroup", "swaggerExample", com.azure.core.util.Context.NONE);
    }
}
