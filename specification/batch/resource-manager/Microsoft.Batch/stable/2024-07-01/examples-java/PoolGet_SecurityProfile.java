
/**
 * Samples for Pool Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/batch/resource-manager/Microsoft.Batch/stable/2024-07-01/examples/PoolGet_SecurityProfile.json
     */
    /**
     * Sample code: GetPool - SecurityProfile.
     * 
     * @param manager Entry point to BatchManager.
     */
    public static void getPoolSecurityProfile(com.azure.resourcemanager.batch.BatchManager manager) {
        manager.pools().getWithResponse("default-azurebatch-japaneast", "sampleacct", "testpool",
            com.azure.core.util.Context.NONE);
    }
}
