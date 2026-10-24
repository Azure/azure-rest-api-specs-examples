
/**
 * Samples for PrivateLinkResources Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-10-01-preview/PrivateLinkResources_Get.json
     */
    /**
     * Sample code: Get a HorizonDB private link resource.
     * 
     * @param manager Entry point to HorizonDbManager.
     */
    public static void getAHorizonDBPrivateLinkResource(com.azure.resourcemanager.horizondb.HorizonDbManager manager) {
        manager.privateLinkResources().getWithResponse("example-resource-group", "example-cluster", "example-cluster",
            com.azure.core.util.Context.NONE);
    }
}
