
/**
 * Samples for StorageContainers Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-06-01/StorageContainers_Delete_MaximumSet_Gen.json
     */
    /**
     * Sample code: StorageContainers_Delete_MaximumSet.
     * 
     * @param manager Entry point to DiscoveryManager.
     */
    public static void storageContainersDeleteMaximumSet(com.azure.resourcemanager.discovery.DiscoveryManager manager) {
        manager.storageContainers().delete("rgdiscovery", "349ff9e95865f956f6", com.azure.core.util.Context.NONE);
    }
}
