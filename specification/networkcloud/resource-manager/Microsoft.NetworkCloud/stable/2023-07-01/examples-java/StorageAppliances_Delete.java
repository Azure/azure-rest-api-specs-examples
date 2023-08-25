/** Samples for StorageAppliances Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/networkcloud/resource-manager/Microsoft.NetworkCloud/stable/2023-07-01/examples/StorageAppliances_Delete.json
     */
    /**
     * Sample code: Delete storage appliance.
     *
     * @param manager Entry point to NetworkCloudManager.
     */
    public static void deleteStorageAppliance(com.azure.resourcemanager.networkcloud.NetworkCloudManager manager) {
        manager
            .storageAppliances()
            .delete("resourceGroupName", "storageApplianceName", com.azure.core.util.Context.NONE);
    }
}
