
/**
 * Samples for SecurityPartnerProviders Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/SecurityPartnerProviderDelete.json
     */
    /**
     * Sample code: Delete Security Partner Provider.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void deleteSecurityPartnerProvider(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getSecurityPartnerProviders().delete("rg1", "securityPartnerProvider",
            com.azure.core.util.Context.NONE);
    }
}
