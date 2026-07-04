
/**
 * Samples for AdminRuleCollections Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/NetworkManagerAdminRuleCollectionDelete.json
     */
    /**
     * Sample code: Deletes an admin rule collection.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void deletesAnAdminRuleCollection(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getAdminRuleCollections().delete("rg1", "testNetworkManager", "myTestSecurityConfig",
            "testRuleCollection", false, com.azure.core.util.Context.NONE);
    }
}
