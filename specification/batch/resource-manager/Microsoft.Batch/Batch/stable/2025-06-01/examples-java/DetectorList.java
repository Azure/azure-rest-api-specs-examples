
/**
 * Samples for BatchAccount ListDetectors.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-06-01/DetectorList.json
     */
    /**
     * Sample code: ListDetectors.
     * 
     * @param manager Entry point to BatchManager.
     */
    public static void listDetectors(com.azure.resourcemanager.batch.BatchManager manager) {
        manager.batchAccounts().listDetectors("default-azurebatch-japaneast", "sampleacct",
            com.azure.core.util.Context.NONE);
    }
}
