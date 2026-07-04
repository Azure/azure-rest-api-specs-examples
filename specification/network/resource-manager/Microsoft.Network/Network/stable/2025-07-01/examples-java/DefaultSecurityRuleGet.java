
/**
 * Samples for DefaultSecurityRules Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/DefaultSecurityRuleGet.json
     */
    /**
     * Sample code: DefaultSecurityRuleGet.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void defaultSecurityRuleGet(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getDefaultSecurityRules().getWithResponse("testrg", "nsg1", "AllowVnetInBound",
            com.azure.core.util.Context.NONE);
    }
}
