
/**
 * Samples for BatchAccount SynchronizeAutoStorageKeys.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-06-01/BatchAccountSynchronizeAutoStorageKeys.json
     */
    /**
     * Sample code: BatchAccountSynchronizeAutoStorageKeys.
     * 
     * @param manager Entry point to BatchManager.
     */
    public static void batchAccountSynchronizeAutoStorageKeys(com.azure.resourcemanager.batch.BatchManager manager) {
        manager.batchAccounts().synchronizeAutoStorageKeysWithResponse("default-azurebatch-japaneast", "sampleacct",
            com.azure.core.util.Context.NONE);
    }
}
