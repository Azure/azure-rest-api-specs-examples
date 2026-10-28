
/**
 * Samples for CustomDomains EnableCustomHttps.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-10-01/CustomDomains_EnableCustomHttpsUsingCDNManagedCertificate.json
     */
    /**
     * Sample code: CustomDomains_EnableCustomHttpsUsingCDNManagedCertificate.
     * 
     * @param manager Entry point to CdnManager.
     */
    public static void
        customDomainsEnableCustomHttpsUsingCDNManagedCertificate(com.azure.resourcemanager.cdn.CdnManager manager) {
        manager.serviceClient().getCustomDomains().enableCustomHttps("RG", "profile1", "endpoint1",
            "www-someDomain-net", null, com.azure.core.util.Context.NONE);
    }
}
