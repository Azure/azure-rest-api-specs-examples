
/**
 * Samples for NetworkSecurityPerimeterAccessRules Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/NspAccessRuleGet.json
     */
    /**
     * Sample code: NspAccessRuleGet.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void nspAccessRuleGet(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getNetworkSecurityPerimeterAccessRules().getWithResponse("rg1", "nsp1", "profile1",
            "accessRule1", com.azure.core.util.Context.NONE);
    }
}
