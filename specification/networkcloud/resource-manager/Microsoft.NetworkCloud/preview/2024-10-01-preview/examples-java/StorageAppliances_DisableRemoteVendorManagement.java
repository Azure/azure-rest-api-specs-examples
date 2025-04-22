
/**
 * Samples for StorageAppliances DisableRemoteVendorManagement.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/networkcloud/resource-manager/Microsoft.NetworkCloud/preview/2024-10-01-preview/examples/
     * StorageAppliances_DisableRemoteVendorManagement.json
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
