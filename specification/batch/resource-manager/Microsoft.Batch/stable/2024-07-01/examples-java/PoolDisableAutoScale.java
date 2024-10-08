
/**
 * Samples for Pool DisableAutoScale.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/batch/resource-manager/Microsoft.Batch/stable/2024-07-01/examples/PoolDisableAutoScale.json
     */
    /**
     * Sample code: Disable AutoScale.
     * 
     * @param manager Entry point to BatchManager.
     */
    public static void disableAutoScale(com.azure.resourcemanager.batch.BatchManager manager) {
        manager.pools().disableAutoScaleWithResponse("default-azurebatch-japaneast", "sampleacct", "testpool",
            com.azure.core.util.Context.NONE);
    }
}
