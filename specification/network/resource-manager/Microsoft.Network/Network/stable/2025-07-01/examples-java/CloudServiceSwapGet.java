
/**
 * Samples for VipSwap Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/CloudServiceSwapGet.json
     */
    /**
     * Sample code: Get swap resource.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void getSwapResource(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getVipSwaps().getWithResponse("rg1", "testCloudService",
            com.azure.core.util.Context.NONE);
    }
}
