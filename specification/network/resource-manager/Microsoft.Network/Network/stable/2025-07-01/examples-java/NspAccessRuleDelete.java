
/**
 * Samples for NetworkSecurityPerimeterAccessRules Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/NspAccessRuleDelete.json
     */
    /**
     * Sample code: NspAccessRulesDelete.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void nspAccessRulesDelete(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getNetworkSecurityPerimeterAccessRules().deleteWithResponse("rg1", "nsp1", "profile1",
            "accessRule1", com.azure.core.util.Context.NONE);
    }
}
