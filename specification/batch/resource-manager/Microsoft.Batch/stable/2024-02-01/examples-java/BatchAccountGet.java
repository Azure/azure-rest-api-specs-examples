
/**
 * Samples for BatchAccount GetByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/batch/resource-manager/Microsoft.Batch/stable/2024-02-01/examples/BatchAccountGet.json
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
