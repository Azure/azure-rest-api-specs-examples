
/**
 * Samples for CustomDomains Delete.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/cdn/resource-manager/Microsoft.Cdn/stable/2024-02-01/examples/CustomDomains_Delete.json
     */
    /**
     * Sample code: CustomDomains_Delete.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void customDomainsDelete(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.cdnProfiles().manager().serviceClient().getCustomDomains().delete("RG", "profile1", "endpoint1",
            "www-someDomain-net", com.azure.core.util.Context.NONE);
    }
}
