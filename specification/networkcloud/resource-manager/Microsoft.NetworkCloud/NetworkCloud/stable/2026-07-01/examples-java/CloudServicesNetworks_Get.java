
/**
 * Samples for CloudServicesNetworks GetByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-07-01/CloudServicesNetworks_Get.json
     */
    /**
     * Sample code: Get cloud services network.
     * 
     * @param manager Entry point to NetworkCloudManager.
     */
    public static void getCloudServicesNetwork(com.azure.resourcemanager.networkcloud.NetworkCloudManager manager) {
        manager.cloudServicesNetworks().getByResourceGroupWithResponse("resourceGroupName", "cloudServicesNetworkName",
            com.azure.core.util.Context.NONE);
    }
}
