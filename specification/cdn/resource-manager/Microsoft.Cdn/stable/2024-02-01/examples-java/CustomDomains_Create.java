
import com.azure.resourcemanager.cdn.models.CustomDomainParameters;

/**
 * Samples for CustomDomains Create.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/cdn/resource-manager/Microsoft.Cdn/stable/2024-02-01/examples/CustomDomains_Create.json
     */
    /**
     * Sample code: CustomDomains_Create.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void customDomainsCreate(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.cdnProfiles().manager().serviceClient().getCustomDomains().create("RG", "profile1", "endpoint1",
            "www-someDomain-net", new CustomDomainParameters().withHostname("www.someDomain.net"),
            com.azure.core.util.Context.NONE);
    }
}
