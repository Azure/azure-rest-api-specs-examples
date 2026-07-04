
/**
 * Samples for VipSwap List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/CloudServiceSwapList.json
     */
    /**
     * Sample code: Get swap resource list.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void getSwapResourceList(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getVipSwaps().listWithResponse("rg1", "testCloudService",
            com.azure.core.util.Context.NONE);
    }
}
