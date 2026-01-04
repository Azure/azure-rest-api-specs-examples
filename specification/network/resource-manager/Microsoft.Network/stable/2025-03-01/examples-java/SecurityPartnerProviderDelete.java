
/**
 * Samples for SecurityPartnerProviders Delete.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/network/resource-manager/Microsoft.Network/stable/2025-03-01/examples/SecurityPartnerProviderDelete
     * .json
     */
    /**
     * Sample code: Delete Security Partner Provider.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void deleteSecurityPartnerProvider(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.networks().manager().serviceClient().getSecurityPartnerProviders().delete("rg1",
            "securityPartnerProvider", com.azure.core.util.Context.NONE);
    }
}
