import com.azure.core.util.Context;

/** Samples for CustomDomains EnableCustomHttps. */
public final class Main {
    /*
     * x-ms-original-file: specification/cdn/resource-manager/Microsoft.Cdn/stable/2021-06-01/examples/CustomDomains_EnableCustomHttpsUsingCDNManagedCertificate.json
     */
    /**
     * Sample code: CustomDomains_EnableCustomHttpsUsingCDNManagedCertificate.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void customDomainsEnableCustomHttpsUsingCDNManagedCertificate(
        com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .cdnProfiles()
            .manager()
            .serviceClient()
            .getCustomDomains()
            .enableCustomHttpsWithResponse("RG", "profile1", "endpoint1", "www-someDomain-net", null, Context.NONE);
    }
}
