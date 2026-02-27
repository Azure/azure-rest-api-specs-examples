
/**
 * Samples for SecurityPartnerProviders ListByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2025-05-01/examples/
     * SecurityPartnerProviderListByResourceGroup.json
     */
    /**
     * Sample code: List all Security Partner Providers for a given resource group.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void
        listAllSecurityPartnerProvidersForAGivenResourceGroup(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.networks().manager().serviceClient().getSecurityPartnerProviders().listByResourceGroup("rg1",
            com.azure.core.util.Context.NONE);
    }
}
