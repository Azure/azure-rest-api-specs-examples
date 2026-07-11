
/**
 * Samples for BatchAccount ListByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-06-01/BatchAccountListByResourceGroup.json
     */
    /**
     * Sample code: BatchAccountListByResourceGroup.
     * 
     * @param manager Entry point to BatchManager.
     */
    public static void batchAccountListByResourceGroup(com.azure.resourcemanager.batch.BatchManager manager) {
        manager.batchAccounts().listByResourceGroup("default-azurebatch-japaneast", com.azure.core.util.Context.NONE);
    }
}
