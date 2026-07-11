
/**
 * Samples for BatchAccount List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-06-01/BatchAccountList.json
     */
    /**
     * Sample code: BatchAccountList.
     * 
     * @param manager Entry point to BatchManager.
     */
    public static void batchAccountList(com.azure.resourcemanager.batch.BatchManager manager) {
        manager.batchAccounts().list(com.azure.core.util.Context.NONE);
    }
}
