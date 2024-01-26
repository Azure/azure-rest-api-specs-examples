
/** Samples for AfdOriginGroups ListResourceUsage. */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/cdn/resource-manager/Microsoft.Cdn/stable/2023-05-01/examples/AFDOriginGroups_ListResourceUsage.
     * json
     */
    /**
     * Sample code: AFDOriginGroups_ListResourceUsage.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void aFDOriginGroupsListResourceUsage(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.cdnProfiles().manager().serviceClient().getAfdOriginGroups().listResourceUsage("RG", "profile1",
            "origingroup1", com.azure.core.util.Context.NONE);
    }
}
