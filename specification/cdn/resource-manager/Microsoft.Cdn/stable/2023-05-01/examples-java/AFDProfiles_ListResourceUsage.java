
/** Samples for AfdProfiles ListResourceUsage. */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/cdn/resource-manager/Microsoft.Cdn/stable/2023-05-01/examples/AFDProfiles_ListResourceUsage.json
     */
    /**
     * Sample code: AFDProfiles_ListResourceUsage.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void aFDProfilesListResourceUsage(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.cdnProfiles().manager().serviceClient().getAfdProfiles().listResourceUsage("RG", "profile1",
            com.azure.core.util.Context.NONE);
    }
}
