
/**
 * Samples for SecurityUserRuleCollections Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-01-01/NetworkManagerSecurityUserRuleCollectionGet.json
     */
    /**
     * Sample code: Gets security user rule collection.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void getsSecurityUserRuleCollection(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getSecurityUserRuleCollections().getWithResponse("rg1", "testNetworkManager",
            "myTestSecurityConfig", "testRuleCollection", com.azure.core.util.Context.NONE);
    }
}
