/** Samples for Profiles GetByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/cdn/resource-manager/Microsoft.Cdn/stable/2023-05-01/examples/Profiles_Get.json
     */
    /**
     * Sample code: Profiles_Get.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void profilesGet(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .cdnProfiles()
            .manager()
            .serviceClient()
            .getProfiles()
            .getByResourceGroupWithResponse("RG", "profile1", com.azure.core.util.Context.NONE);
    }
}
