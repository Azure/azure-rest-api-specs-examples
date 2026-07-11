
/**
 * Samples for BatchAccount GetByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-06-01/PrivateBatchAccountGet.json
     */
    /**
     * Sample code: PrivateBatchAccountGet.
     * 
     * @param manager Entry point to BatchManager.
     */
    public static void privateBatchAccountGet(com.azure.resourcemanager.batch.BatchManager manager) {
        manager.batchAccounts().getByResourceGroupWithResponse("default-azurebatch-japaneast", "sampleacct",
            com.azure.core.util.Context.NONE);
    }
}
