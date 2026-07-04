
/**
 * Samples for SecurityPartnerProviders List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/SecurityPartnerProviderListBySubscription.json
     */
    /**
     * Sample code: List all Security Partner Providers for a given subscription.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void
        listAllSecurityPartnerProvidersForAGivenSubscription(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getSecurityPartnerProviders().list(com.azure.core.util.Context.NONE);
    }
}
