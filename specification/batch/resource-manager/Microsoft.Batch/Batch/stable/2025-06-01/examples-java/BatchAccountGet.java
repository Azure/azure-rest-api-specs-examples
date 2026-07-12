
/**
 * Samples for BatchAccount GetByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-06-01/BatchAccountGet.json
     */
    /**
     * Sample code: BatchAccountGet.
     * 
     * @param manager Entry point to BatchManager.
     */
    public static void batchAccountGet(com.azure.resourcemanager.batch.BatchManager manager) {
        manager.batchAccounts().getByResourceGroupWithResponse("default-azurebatch-japaneast", "sampleacct",
            com.azure.core.util.Context.NONE);
    }
}
