import com.azure.core.util.Context;

/** Samples for SecurityPartnerProviders GetByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2021-05-01/examples/SecurityPartnerProviderGet.json
     */
    /**
     * Sample code: Get Security Partner Provider.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getSecurityPartnerProvider(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .networks()
            .manager()
            .serviceClient()
            .getSecurityPartnerProviders()
            .getByResourceGroupWithResponse("rg1", "securityPartnerProvider", Context.NONE);
    }
}
