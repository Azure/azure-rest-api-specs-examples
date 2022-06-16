import com.azure.core.util.Context;

/** Samples for AfdCustomDomains ListByProfile. */
public final class Main {
    /*
     * x-ms-original-file: specification/cdn/resource-manager/Microsoft.Cdn/stable/2021-06-01/examples/AFDCustomDomains_ListByProfile.json
     */
    /**
     * Sample code: AFDCustomDomains_ListByProfile.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void aFDCustomDomainsListByProfile(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .cdnProfiles()
            .manager()
            .serviceClient()
            .getAfdCustomDomains()
            .listByProfile("RG", "profile1", Context.NONE);
    }
}
