/** Samples for DppResourceGuardProxy Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/dataprotection/resource-manager/Microsoft.DataProtection/stable/2023-01-01/examples/ResourceGuardProxyCRUD/DeleteResourceGuardProxy.json
     */
    /**
     * Sample code: Delete ResourceGuardProxy.
     *
     * @param manager Entry point to DataProtectionManager.
     */
    public static void deleteResourceGuardProxy(
        com.azure.resourcemanager.dataprotection.DataProtectionManager manager) {
        manager
            .dppResourceGuardProxies()
            .deleteWithResponse(
                "SampleResourceGroup", "sampleVault", "swaggerExample", com.azure.core.util.Context.NONE);
    }
}
