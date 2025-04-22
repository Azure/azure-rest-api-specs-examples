
/**
 * Samples for StandbyVirtualMachinePoolRuntimeViews ListByStandbyPool.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-03-01/StandbyVirtualMachinePoolRuntimeViews_ListByStandbyPool.json
     */
    /**
     * Sample code: StandbyVirtualMachinePoolRuntimeViews_ListByStandbyPool.
     * 
     * @param manager Entry point to StandbyPoolManager.
     */
    public static void standbyVirtualMachinePoolRuntimeViewsListByStandbyPool(
        com.azure.resourcemanager.standbypool.StandbyPoolManager manager) {
        manager.standbyVirtualMachinePoolRuntimeViews().listByStandbyPool("rgstandbypool", "pool",
            com.azure.core.util.Context.NONE);
    }
}
