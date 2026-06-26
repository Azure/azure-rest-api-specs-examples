
/**
 * Samples for StorageAppliances Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-05-01-preview/StorageAppliances_Delete.json
     */
    /**
     * Sample code: Delete storage appliance.
     * 
     * @param manager Entry point to NetworkCloudManager.
     */
    public static void deleteStorageAppliance(com.azure.resourcemanager.networkcloud.NetworkCloudManager manager) {
        manager.storageAppliances().delete("resourceGroupName", "storageApplianceName", null, null,
            com.azure.core.util.Context.NONE);
    }
}
