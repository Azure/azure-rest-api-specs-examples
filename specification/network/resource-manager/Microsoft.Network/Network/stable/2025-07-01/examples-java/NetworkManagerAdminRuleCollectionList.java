
/**
 * Samples for AdminRuleCollections List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/NetworkManagerAdminRuleCollectionList.json
     */
    /**
     * Sample code: List security admin rule collections.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void listSecurityAdminRuleCollections(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getAdminRuleCollections().list("rg1", "testNetworkManager", "myTestSecurityConfig",
            null, null, com.azure.core.util.Context.NONE);
    }
}
