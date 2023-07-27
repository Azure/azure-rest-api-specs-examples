/** Samples for Pool ListByBatchAccount. */
public final class Main {
    /*
     * x-ms-original-file: specification/batch/resource-manager/Microsoft.Batch/stable/2023-05-01/examples/PoolList.json
     */
    /**
     * Sample code: ListPool.
     *
     * @param manager Entry point to BatchManager.
     */
    public static void listPool(com.azure.resourcemanager.batch.BatchManager manager) {
        manager
            .pools()
            .listByBatchAccount(
                "default-azurebatch-japaneast", "sampleacct", null, null, null, com.azure.core.util.Context.NONE);
    }
}
