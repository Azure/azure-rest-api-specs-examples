
/**
 * Samples for Routes Get.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/cdn/resource-manager/Microsoft.Cdn/stable/2024-02-01/examples/Routes_Get.json
     */
    /**
     * Sample code: Routes_Get.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void routesGet(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.cdnProfiles().manager().serviceClient().getRoutes().getWithResponse("RG", "profile1", "endpoint1",
            "route1", com.azure.core.util.Context.NONE);
    }
}
