
/**
 * Samples for StorageAppliances ListByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/networkcloud/resource-manager/Microsoft.NetworkCloud/stable/2025-09-01/examples/
     * StorageAppliances_ListByResourceGroup.json
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
