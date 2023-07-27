/** Samples for BatchAccount GetKeys. */
public final class Main {
    /*
     * x-ms-original-file: specification/batch/resource-manager/Microsoft.Batch/stable/2023-05-01/examples/BatchAccountGetKeys.json
     */
    /**
     * Sample code: BatchAccountGetKeys.
     *
     * @param manager Entry point to BatchManager.
     */
    public static void batchAccountGetKeys(com.azure.resourcemanager.batch.BatchManager manager) {
        manager
            .batchAccounts()
            .getKeysWithResponse("default-azurebatch-japaneast", "sampleacct", com.azure.core.util.Context.NONE);
    }
}
