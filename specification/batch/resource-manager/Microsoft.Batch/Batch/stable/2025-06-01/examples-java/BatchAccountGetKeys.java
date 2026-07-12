
/**
 * Samples for BatchAccount GetKeys.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-06-01/BatchAccountGetKeys.json
     */
    /**
     * Sample code: BatchAccountGetKeys.
     * 
     * @param manager Entry point to BatchManager.
     */
    public static void batchAccountGetKeys(com.azure.resourcemanager.batch.BatchManager manager) {
        manager.batchAccounts().getKeysWithResponse("default-azurebatch-japaneast", "sampleacct",
            com.azure.core.util.Context.NONE);
    }
}
