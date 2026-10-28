
/**
 * Samples for CustomDomains DisableCustomHttps.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-10-01/CustomDomains_DisableCustomHttps.json
     */
    /**
     * Sample code: CustomDomains_DisableCustomHttps.
     * 
     * @param manager Entry point to CdnManager.
     */
    public static void customDomainsDisableCustomHttps(com.azure.resourcemanager.cdn.CdnManager manager) {
        manager.serviceClient().getCustomDomains().disableCustomHttps("RG", "profile1", "endpoint1",
            "www-someDomain-net", com.azure.core.util.Context.NONE);
    }
}
