import com.azure.core.util.Context;

/** Samples for CustomDomains DisableCustomHttps. */
public final class Main {
    /*
     * x-ms-original-file: specification/cdn/resource-manager/Microsoft.Cdn/stable/2021-06-01/examples/CustomDomains_DisableCustomHttps.json
     */
    /**
     * Sample code: CustomDomains_DisableCustomHttps.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void customDomainsDisableCustomHttps(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .cdnProfiles()
            .manager()
            .serviceClient()
            .getCustomDomains()
            .disableCustomHttpsWithResponse("RG", "profile1", "endpoint1", "www-someDomain-net", Context.NONE);
    }
}
