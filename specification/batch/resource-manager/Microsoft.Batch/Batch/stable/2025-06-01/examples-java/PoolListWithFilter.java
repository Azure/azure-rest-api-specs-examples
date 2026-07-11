
/**
 * Samples for Pool ListByBatchAccount.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-06-01/PoolListWithFilter.json
     */
    /**
     * Sample code: ListPoolWithFilter.
     * 
     * @param manager Entry point to BatchManager.
     */
    public static void listPoolWithFilter(com.azure.resourcemanager.batch.BatchManager manager) {
        manager.pools().listByBatchAccount("default-azurebatch-japaneast", "sampleacct", 50,
            "properties/allocationState,properties/provisioningStateTransitionTime,properties/currentDedicatedNodes,properties/currentLowPriorityNodes",
            "startswith(name, 'po') or (properties/allocationState eq 'Steady' and properties/provisioningStateTransitionTime lt datetime'2017-02-02')",
            com.azure.core.util.Context.NONE);
    }
}
