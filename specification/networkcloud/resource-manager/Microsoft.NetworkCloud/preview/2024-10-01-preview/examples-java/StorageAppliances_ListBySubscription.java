
/**
 * Samples for StorageAppliances List.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/networkcloud/resource-manager/Microsoft.NetworkCloud/preview/2024-10-01-preview/examples/
     * StorageAppliances_ListBySubscription.json
     */
    /**
     * Sample code: List storage appliances for subscription.
     * 
     * @param manager Entry point to NetworkCloudManager.
     */
    public static void
        listStorageAppliancesForSubscription(com.azure.resourcemanager.networkcloud.NetworkCloudManager manager) {
        manager.storageAppliances().list(com.azure.core.util.Context.NONE);
    }
}
