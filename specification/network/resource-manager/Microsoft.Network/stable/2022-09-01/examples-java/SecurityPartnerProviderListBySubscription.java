/** Samples for SecurityPartnerProviders List. */
public final class Main {
    /*
     * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2022-09-01/examples/SecurityPartnerProviderListBySubscription.json
     */
    /**
     * Sample code: List all Security Partner Providers for a given subscription.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void listAllSecurityPartnerProvidersForAGivenSubscription(
        com.azure.resourcemanager.AzureResourceManager azure) {
        azure.networks().manager().serviceClient().getSecurityPartnerProviders().list(com.azure.core.util.Context.NONE);
    }
}
