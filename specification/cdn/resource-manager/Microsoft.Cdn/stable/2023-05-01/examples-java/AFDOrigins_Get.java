/** Samples for AfdOrigins Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/cdn/resource-manager/Microsoft.Cdn/stable/2023-05-01/examples/AFDOrigins_Get.json
     */
    /**
     * Sample code: AFDOrigins_Get.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void aFDOriginsGet(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .cdnProfiles()
            .manager()
            .serviceClient()
            .getAfdOrigins()
            .getWithResponse("RG", "profile1", "origingroup1", "origin1", com.azure.core.util.Context.NONE);
    }
}
