
/**
 * Samples for StorageDiscoveryWorkspaces List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-06-01-preview/StorageDiscoveryWorkspaces_ListBySubscription.json
     */
    /**
     * Sample code: List StorageDiscoveryWorkspaces by Subscription.
     * 
     * @param manager Entry point to StorageDiscoveryManager.
     */
    public static void listStorageDiscoveryWorkspacesBySubscription(
        com.azure.resourcemanager.storagediscovery.StorageDiscoveryManager manager) {
        manager.storageDiscoveryWorkspaces().list(com.azure.core.util.Context.NONE);
    }
}
