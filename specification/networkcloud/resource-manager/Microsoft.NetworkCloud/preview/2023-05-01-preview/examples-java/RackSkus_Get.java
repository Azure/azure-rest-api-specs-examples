/** Samples for RackSkus Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/networkcloud/resource-manager/Microsoft.NetworkCloud/preview/2023-05-01-preview/examples/RackSkus_Get.json
     */
    /**
     * Sample code: Get rack SKU resource.
     *
     * @param manager Entry point to NetworkCloudManager.
     */
    public static void getRackSKUResource(com.azure.resourcemanager.networkcloud.NetworkCloudManager manager) {
        manager.rackSkus().getWithResponse("rackSkuName", com.azure.core.util.Context.NONE);
    }
}
