import com.azure.core.util.Context;

/** Samples for BatchAccount GetByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/batch/resource-manager/Microsoft.Batch/stable/2022-01-01/examples/PrivateBatchAccountGet.json
     */
    /**
     * Sample code: PrivateBatchAccountGet.
     *
     * @param manager Entry point to BatchManager.
     */
    public static void privateBatchAccountGet(com.azure.resourcemanager.batch.BatchManager manager) {
        manager
            .batchAccounts()
            .getByResourceGroupWithResponse("default-azurebatch-japaneast", "sampleacct", Context.NONE);
    }
}
