
/**
 * Samples for StandbyContainerGroupPoolRuntimeViews ListByStandbyPool.
 */
public final class Main {
    /*
     * x-ms-original-file: 2024-03-01/StandbyContainerGroupPoolRuntimeViews_ListByStandbyPool.json
     */
    /**
     * Sample code: StandbyContainerGroupPoolRuntimeViews_ListByStandbyPool.
     * 
     * @param manager Entry point to StandbyPoolManager.
     */
    public static void standbyContainerGroupPoolRuntimeViewsListByStandbyPool(
        com.azure.resourcemanager.standbypool.StandbyPoolManager manager) {
        manager.standbyContainerGroupPoolRuntimeViews().listByStandbyPool("rgstandbypool", "pool",
            com.azure.core.util.Context.NONE);
    }
}
