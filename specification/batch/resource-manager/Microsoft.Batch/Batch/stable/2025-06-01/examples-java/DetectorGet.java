
/**
 * Samples for BatchAccount GetDetector.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-06-01/DetectorGet.json
     */
    /**
     * Sample code: GetDetector.
     * 
     * @param manager Entry point to BatchManager.
     */
    public static void getDetector(com.azure.resourcemanager.batch.BatchManager manager) {
        manager.batchAccounts().getDetectorWithResponse("default-azurebatch-japaneast", "sampleacct", "poolsAndNodes",
            com.azure.core.util.Context.NONE);
    }
}
