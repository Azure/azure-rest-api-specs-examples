
/**
 * Samples for SecurityUserRuleCollections Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/NetworkManagerSecurityUserRuleCollectionDelete.json
     */
    /**
     * Sample code: Deletes a Security User Rule collection.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void deletesASecurityUserRuleCollection(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getSecurityUserRuleCollections().delete("rg1", "testNetworkManager",
            "myTestSecurityConfig", "testRuleCollection", false, com.azure.core.util.Context.NONE);
    }
}
