/** Samples for Volumes GetByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/networkcloud/resource-manager/Microsoft.NetworkCloud/preview/2023-05-01-preview/examples/Volumes_Get.json
     */
    /**
     * Sample code: Get volume.
     *
     * @param manager Entry point to NetworkCloudManager.
     */
    public static void getVolume(com.azure.resourcemanager.networkcloud.NetworkCloudManager manager) {
        manager
            .volumes()
            .getByResourceGroupWithResponse("resourceGroupName", "volumeName", com.azure.core.util.Context.NONE);
    }
}
