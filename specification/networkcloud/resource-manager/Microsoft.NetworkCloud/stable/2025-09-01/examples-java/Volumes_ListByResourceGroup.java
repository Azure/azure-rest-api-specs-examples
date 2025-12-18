
/**
 * Samples for Volumes ListByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/networkcloud/resource-manager/Microsoft.NetworkCloud/stable/2025-09-01/examples/
     * Volumes_ListByResourceGroup.json
     */
    /**
     * Sample code: List volumes for resource group.
     * 
     * @param manager Entry point to NetworkCloudManager.
     */
    public static void listVolumesForResourceGroup(com.azure.resourcemanager.networkcloud.NetworkCloudManager manager) {
        manager.volumes().listByResourceGroup("resourceGroupName", null, null, com.azure.core.util.Context.NONE);
    }
}
