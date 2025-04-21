
/**
 * Samples for StandbyContainerGroupPools List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-03-01/StandbyContainerGroupPools_ListBySubscription.json
     */
    /**
     * Sample code: StandbyContainerGroupPools_ListBySubscription.
     * 
     * @param manager Entry point to StandbyPoolManager.
     */
    public static void
        standbyContainerGroupPoolsListBySubscription(com.azure.resourcemanager.standbypool.StandbyPoolManager manager) {
        manager.standbyContainerGroupPools().list(com.azure.core.util.Context.NONE);
    }
}
