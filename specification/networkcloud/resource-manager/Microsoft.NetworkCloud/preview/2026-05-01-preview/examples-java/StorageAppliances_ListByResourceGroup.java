
/**
 * Samples for StorageAppliances ListByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-05-01-preview/StorageAppliances_ListByResourceGroup.json
     */
    /**
     * Sample code: List storage appliances for resource group.
     * 
     * @param manager Entry point to NetworkCloudManager.
     */
    public static void
        listStorageAppliancesForResourceGroup(com.azure.resourcemanager.networkcloud.NetworkCloudManager manager) {
        manager.storageAppliances().listByResourceGroup("resourceGroupName", null, null,
            com.azure.core.util.Context.NONE);
    }
}
