
/**
 * Samples for Endpoints Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-10-01-preview/Endpoints_Get.json
     */
    /**
     * Sample code: Get a HorizonDB endpoint.
     * 
     * @param manager Entry point to HorizonDbManager.
     */
    public static void getAHorizonDBEndpoint(com.azure.resourcemanager.horizondb.HorizonDbManager manager) {
        manager.endpoints().getWithResponse("example-resource-group", "example-cluster", "example-cluster",
            com.azure.core.util.Context.NONE);
    }
}
