
/**
 * Samples for StorageAppliances List.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/networkcloud/resource-manager/Microsoft.NetworkCloud/stable/2025-09-01/examples/
     * StorageAppliances_ListBySubscription.json
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
