
/**
 * Samples for SecurityPartnerProviders GetByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/SecurityPartnerProviderGet.json
     */
    /**
     * Sample code: Get Security Partner Provider.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void getSecurityPartnerProvider(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getSecurityPartnerProviders().getByResourceGroupWithResponse("rg1",
            "securityPartnerProvider", com.azure.core.util.Context.NONE);
    }
}
