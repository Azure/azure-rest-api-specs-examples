
/**
 * Samples for Pool Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-06-01/PoolGet_VirtualMachineConfiguration_ServiceArtifactReference.json
     */
    /**
     * Sample code: GetPool - VirtualMachineConfiguration ServiceArtifactReference.
     * 
     * @param manager Entry point to BatchManager.
     */
    public static void getPoolVirtualMachineConfigurationServiceArtifactReference(
        com.azure.resourcemanager.batch.BatchManager manager) {
        manager.pools().getWithResponse("default-azurebatch-japaneast", "sampleacct", "testpool",
            com.azure.core.util.Context.NONE);
    }
}
