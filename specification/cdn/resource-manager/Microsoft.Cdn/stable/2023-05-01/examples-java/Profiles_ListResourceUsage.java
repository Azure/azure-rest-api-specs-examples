/** Samples for Profiles ListResourceUsage. */
public final class Main {
    /*
     * x-ms-original-file: specification/cdn/resource-manager/Microsoft.Cdn/stable/2023-05-01/examples/Profiles_ListResourceUsage.json
     */
    /**
     * Sample code: Profiles_ListResourceUsage.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void profilesListResourceUsage(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .cdnProfiles()
            .manager()
            .serviceClient()
            .getProfiles()
            .listResourceUsage("RG", "profile1", com.azure.core.util.Context.NONE);
    }
}
