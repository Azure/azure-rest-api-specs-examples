
/**
 * Samples for VnetConnections Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-07-01/VnetConnections_Get.json
     */
    /**
     * Sample code: Get a VnetConnection.
     * 
     * @param manager Entry point to ContainerAppsApiManager.
     */
    public static void getAVnetConnection(com.azure.resourcemanager.appcontainers.ContainerAppsApiManager manager) {
        manager.vnetConnections().getWithResponse("myRg", "testgroup", "myVnetConnection",
            com.azure.core.util.Context.NONE);
    }
}
