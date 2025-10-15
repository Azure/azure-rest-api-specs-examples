
/**
 * Samples for DppResourceGuardProxy List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/ResourceGuardProxyCRUD/ListResourceGuardProxy.json
     */
    /**
     * Sample code: Get ResourceGuardProxies.
     * 
     * @param manager Entry point to DataProtectionManager.
     */
    public static void getResourceGuardProxies(com.azure.resourcemanager.dataprotection.DataProtectionManager manager) {
        manager.dppResourceGuardProxies().list("SampleResourceGroup", "sampleVault", com.azure.core.util.Context.NONE);
    }
}
