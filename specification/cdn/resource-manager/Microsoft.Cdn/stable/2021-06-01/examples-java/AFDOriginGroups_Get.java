import com.azure.core.util.Context;

/** Samples for AfdOriginGroups Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/cdn/resource-manager/Microsoft.Cdn/stable/2021-06-01/examples/AFDOriginGroups_Get.json
     */
    /**
     * Sample code: AFDOriginGroups_Get.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void aFDOriginGroupsGet(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .cdnProfiles()
            .manager()
            .serviceClient()
            .getAfdOriginGroups()
            .getWithResponse("RG", "profile1", "origingroup1", Context.NONE);
    }
}
