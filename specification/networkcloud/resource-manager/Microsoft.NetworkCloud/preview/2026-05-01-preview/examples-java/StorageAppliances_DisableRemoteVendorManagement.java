
/**
 * Samples for StorageAppliances DisableRemoteVendorManagement.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-05-01-preview/StorageAppliances_DisableRemoteVendorManagement.json
     */
    /**
     * Sample code: Turn off remote vendor management for storage appliance.
     * 
     * @param manager Entry point to NetworkCloudManager.
     */
    public static void turnOffRemoteVendorManagementForStorageAppliance(
        com.azure.resourcemanager.networkcloud.NetworkCloudManager manager) {
        manager.storageAppliances().disableRemoteVendorManagement("resourceGroupName", "storageApplianceName",
            com.azure.core.util.Context.NONE);
    }
}
