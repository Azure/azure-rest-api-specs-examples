/** Samples for BatchAccount ListByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/batch/resource-manager/Microsoft.Batch/stable/2023-05-01/examples/BatchAccountListByResourceGroup.json
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
