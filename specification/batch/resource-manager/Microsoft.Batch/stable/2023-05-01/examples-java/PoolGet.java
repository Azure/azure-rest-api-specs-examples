/** Samples for Pool Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/batch/resource-manager/Microsoft.Batch/stable/2023-05-01/examples/PoolGet.json
     */
    /**
     * Sample code: GetPool.
     *
     * @param manager Entry point to BatchManager.
     */
    public static void getPool(com.azure.resourcemanager.batch.BatchManager manager) {
        manager
            .pools()
            .getWithResponse(
                "default-azurebatch-japaneast", "sampleacct", "testpool", com.azure.core.util.Context.NONE);
    }
}
