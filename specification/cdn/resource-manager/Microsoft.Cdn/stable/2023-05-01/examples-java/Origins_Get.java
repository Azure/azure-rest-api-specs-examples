/** Samples for Origins Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/cdn/resource-manager/Microsoft.Cdn/stable/2023-05-01/examples/Origins_Get.json
     */
    /**
     * Sample code: Origins_Get.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void originsGet(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .cdnProfiles()
            .manager()
            .serviceClient()
            .getOrigins()
            .getWithResponse("RG", "profile1", "endpoint1", "www-someDomain-net", com.azure.core.util.Context.NONE);
    }
}
