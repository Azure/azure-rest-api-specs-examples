import com.azure.core.util.Context;

/** Samples for CustomDomains ListByEndpoint. */
public final class Main {
    /*
     * x-ms-original-file: specification/cdn/resource-manager/Microsoft.Cdn/stable/2021-06-01/examples/CustomDomains_ListByEndpoint.json
     */
    /**
     * Sample code: CustomDomains_ListByEndpoint.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void customDomainsListByEndpoint(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .cdnProfiles()
            .manager()
            .serviceClient()
            .getCustomDomains()
            .listByEndpoint("RG", "profile1", "endpoint1", Context.NONE);
    }
}
