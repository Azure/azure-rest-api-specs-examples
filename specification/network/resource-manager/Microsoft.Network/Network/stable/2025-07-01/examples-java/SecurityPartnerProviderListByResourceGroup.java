
/**
 * Samples for SecurityPartnerProviders ListByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/SecurityPartnerProviderListByResourceGroup.json
     */
    /**
     * Sample code: List all Security Partner Providers for a given resource group.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void listAllSecurityPartnerProvidersForAGivenResourceGroup(
        com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getSecurityPartnerProviders().listByResourceGroup("rg1",
            com.azure.core.util.Context.NONE);
    }
}
