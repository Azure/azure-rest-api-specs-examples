
/**
 * Samples for AfdCustomDomains RefreshValidationToken.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/cdn/resource-manager/Microsoft.Cdn/stable/2024-02-01/examples/
     * AFDCustomDomains_RefreshValidationToken.json
     */
    /**
     * Sample code: AFDCustomDomains_Delete.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void aFDCustomDomainsDelete(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.cdnProfiles().manager().serviceClient().getAfdCustomDomains().refreshValidationToken("RG", "profile1",
            "domain1", com.azure.core.util.Context.NONE);
    }
}
