
/**
 * Samples for Endpoints Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-10-01-preview/Endpoints_GetReadOnly.json
     */
    /**
     * Sample code: Get a read-only HorizonDB pool endpoint.
     * 
     * @param manager Entry point to HorizonDbManager.
     */
    public static void getAReadOnlyHorizonDBPoolEndpoint(com.azure.resourcemanager.horizondb.HorizonDbManager manager) {
        manager.endpoints().getWithResponse("example-resource-group", "example-cluster", "example-pool-01",
            com.azure.core.util.Context.NONE);
    }
}
