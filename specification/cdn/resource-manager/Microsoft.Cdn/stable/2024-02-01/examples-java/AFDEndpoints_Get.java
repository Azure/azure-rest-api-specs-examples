
/**
 * Samples for AfdEndpoints Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/cdn/resource-manager/Microsoft.Cdn/stable/2024-02-01/examples/AFDEndpoints_Get.json
     */
    /**
     * Sample code: AFDEndpoints_Get.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void aFDEndpointsGet(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.cdnProfiles().manager().serviceClient().getAfdEndpoints().getWithResponse("RG", "profile1", "endpoint1",
            com.azure.core.util.Context.NONE);
    }
}
