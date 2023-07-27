/** Samples for BatchAccount ListDetectors. */
public final class Main {
    /*
     * x-ms-original-file: specification/batch/resource-manager/Microsoft.Batch/stable/2023-05-01/examples/DetectorList.json
     */
    /**
     * Sample code: ListDetectors.
     *
     * @param manager Entry point to BatchManager.
     */
    public static void listDetectors(com.azure.resourcemanager.batch.BatchManager manager) {
        manager
            .batchAccounts()
            .listDetectors("default-azurebatch-japaneast", "sampleacct", com.azure.core.util.Context.NONE);
    }
}
