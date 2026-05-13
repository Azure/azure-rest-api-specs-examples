
/**
 * Samples for DppResourceGuardProxy Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-01/ResourceGuardProxyCRUD/DeleteResourceGuardProxy.json
     */
    /**
     * Sample code: Delete ResourceGuardProxy.
     * 
     * @param manager Entry point to DataProtectionManager.
     */
    public static void
        deleteResourceGuardProxy(com.azure.resourcemanager.dataprotection.DataProtectionManager manager) {
        manager.dppResourceGuardProxies().deleteWithResponse("SampleResourceGroup", "sampleVault", "swaggerExample",
            com.azure.core.util.Context.NONE);
    }
}
