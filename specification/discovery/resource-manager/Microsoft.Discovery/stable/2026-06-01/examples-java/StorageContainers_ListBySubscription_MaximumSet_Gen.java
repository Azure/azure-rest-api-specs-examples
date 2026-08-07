
/**
 * Samples for StorageContainers List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-06-01/StorageContainers_ListBySubscription_MaximumSet_Gen.json
     */
    /**
     * Sample code: StorageContainers_ListBySubscription_MaximumSet.
     * 
     * @param manager Entry point to DiscoveryManager.
     */
    public static void
        storageContainersListBySubscriptionMaximumSet(com.azure.resourcemanager.discovery.DiscoveryManager manager) {
        manager.storageContainers().list(com.azure.core.util.Context.NONE);
    }
}
