
/**
 * Samples for Pool Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-06-01/PoolGet_VirtualMachineConfiguration_Extensions.json
     */
    /**
     * Sample code: GetPool - VirtualMachineConfiguration Extensions.
     * 
     * @param manager Entry point to BatchManager.
     */
    public static void
        getPoolVirtualMachineConfigurationExtensions(com.azure.resourcemanager.batch.BatchManager manager) {
        manager.pools().getWithResponse("default-azurebatch-japaneast", "sampleacct", "testpool",
            com.azure.core.util.Context.NONE);
    }
}
