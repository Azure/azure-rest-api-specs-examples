
/**
 * Samples for SecurityUserRuleCollections List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/NetworkManagerSecurityUserRuleCollectionList.json
     */
    /**
     * Sample code: List rule collections in a security configuration.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void
        listRuleCollectionsInASecurityConfiguration(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getSecurityUserRuleCollections().list("rg1", "testNetworkManager",
            "myTestSecurityConfig", null, null, com.azure.core.util.Context.NONE);
    }
}
