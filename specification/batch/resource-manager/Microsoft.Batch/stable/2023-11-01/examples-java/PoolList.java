/** Samples for Pool ListByBatchAccount. */
public final class Main {
    /*
     * x-ms-original-file: specification/batch/resource-manager/Microsoft.Batch/stable/2023-11-01/examples/PoolList.json
     */
    /**
     * Sample code: ListPool.
     *
     * @param manager Entry point to BatchManager.
     */
    public static void listPool(com.azure.resourcemanager.batch.BatchManager manager) {
        manager
            .pools()
            .listByBatchAccount(
                "default-azurebatch-japaneast", "sampleacct", null, null, null, com.azure.core.util.Context.NONE);
    }

    /*
     * x-ms-original-file:
     * specification/batch/resource-manager/Microsoft.Batch/stable/2023-11-01/examples/PoolListWithFilter.json
     */
    /**
     * Sample code: ListPoolWithFilter.
     *
     * @param manager Entry point to BatchManager.
     */
    public static void listPoolWithFilter(com.azure.resourcemanager.batch.BatchManager manager) {
        manager
            .pools()
            .listByBatchAccount(
                "default-azurebatch-japaneast",
                "sampleacct",
                50,
                "properties/allocationState,properties/provisioningStateTransitionTime,properties/currentDedicatedNodes,properties/currentLowPriorityNodes",
                "startswith(name, 'po') or (properties/allocationState eq 'Steady' and"
                    + " properties/provisioningStateTransitionTime lt datetime'2017-02-02')",
                com.azure.core.util.Context.NONE);
    }
}
