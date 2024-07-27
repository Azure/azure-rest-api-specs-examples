
/**
 * Samples for SecurityPartnerProviders GetByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/network/resource-manager/Microsoft.Network/stable/2024-01-01/examples/SecurityPartnerProviderGet.
     * json
     */
    /**
     * Sample code: Get Security Partner Provider.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getSecurityPartnerProvider(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.networks().manager().serviceClient().getSecurityPartnerProviders().getByResourceGroupWithResponse("rg1",
            "securityPartnerProvider", com.azure.core.util.Context.NONE);
    }
}
