
/**
 * Samples for DppResourceGuardProxy Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/dataprotection/resource-manager/Microsoft.DataProtection/stable/2025-07-01/examples/
     * ResourceGuardProxyCRUD/GetResourceGuardProxy.json
     */
    /**
     * Sample code: Get ResourceGuardProxy.
     * 
     * @param manager Entry point to DataProtectionManager.
     */
    public static void getResourceGuardProxy(com.azure.resourcemanager.dataprotection.DataProtectionManager manager) {
        manager.dppResourceGuardProxies().getWithResponse("SampleResourceGroup", "sampleVault", "swaggerExample",
            com.azure.core.util.Context.NONE);
    }
}
