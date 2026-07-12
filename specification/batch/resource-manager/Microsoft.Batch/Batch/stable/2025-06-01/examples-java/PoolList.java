
/**
 * Samples for Pool ListByBatchAccount.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-06-01/PoolList.json
     */
    /**
     * Sample code: ListPool.
     * 
     * @param manager Entry point to BatchManager.
     */
    public static void listPool(com.azure.resourcemanager.batch.BatchManager manager) {
        manager.pools().listByBatchAccount("default-azurebatch-japaneast", "sampleacct", null, null, null,
            com.azure.core.util.Context.NONE);
    }
}
