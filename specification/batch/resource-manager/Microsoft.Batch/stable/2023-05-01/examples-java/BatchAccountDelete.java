/** Samples for BatchAccount Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/batch/resource-manager/Microsoft.Batch/stable/2023-05-01/examples/BatchAccountDelete.json
     */
    /**
     * Sample code: BatchAccountDelete.
     *
     * @param manager Entry point to BatchManager.
     */
    public static void batchAccountDelete(com.azure.resourcemanager.batch.BatchManager manager) {
        manager.batchAccounts().delete("default-azurebatch-japaneast", "sampleacct", com.azure.core.util.Context.NONE);
    }
}
