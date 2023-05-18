/** Samples for Volumes List. */
public final class Main {
    /*
     * x-ms-original-file: specification/networkcloud/resource-manager/Microsoft.NetworkCloud/preview/2022-12-12-preview/examples/Volumes_ListBySubscription.json
     */
    /**
     * Sample code: List volume for subscription.
     *
     * @param manager Entry point to NetworkCloudManager.
     */
    public static void listVolumeForSubscription(com.azure.resourcemanager.networkcloud.NetworkCloudManager manager) {
        manager.volumes().list(com.azure.core.util.Context.NONE);
    }
}
