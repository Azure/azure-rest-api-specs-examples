
/**
 * Samples for StorageAppliances List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-05-01-preview/StorageAppliances_ListBySubscription.json
     */
    /**
     * Sample code: List storage appliances for subscription.
     * 
     * @param manager Entry point to NetworkCloudManager.
     */
    public static void
        listStorageAppliancesForSubscription(com.azure.resourcemanager.networkcloud.NetworkCloudManager manager) {
        manager.storageAppliances().list(null, null, com.azure.core.util.Context.NONE);
    }
}
