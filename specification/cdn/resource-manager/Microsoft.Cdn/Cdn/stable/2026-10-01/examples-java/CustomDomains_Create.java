
import com.azure.resourcemanager.cdn.models.CustomDomainParameters;

/**
 * Samples for CustomDomains Create.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-10-01/CustomDomains_Create.json
     */
    /**
     * Sample code: CustomDomains_Create.
     * 
     * @param manager Entry point to CdnManager.
     */
    public static void customDomainsCreate(com.azure.resourcemanager.cdn.CdnManager manager) {
        manager.serviceClient().getCustomDomains().create("RG", "profile1", "endpoint1", "www-someDomain-net",
            new CustomDomainParameters().withHostname("www.someDomain.net"), com.azure.core.util.Context.NONE);
    }
}
