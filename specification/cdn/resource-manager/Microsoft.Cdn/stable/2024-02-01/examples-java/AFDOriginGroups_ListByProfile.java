
/**
 * Samples for AfdOriginGroups ListByProfile.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/cdn/resource-manager/Microsoft.Cdn/stable/2024-02-01/examples/AFDOriginGroups_ListByProfile.json
     */
    /**
     * Sample code: AFDOriginGroups_ListByProfile.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void aFDOriginGroupsListByProfile(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.cdnProfiles().manager().serviceClient().getAfdOriginGroups().listByProfile("RG", "profile1",
            com.azure.core.util.Context.NONE);
    }
}
