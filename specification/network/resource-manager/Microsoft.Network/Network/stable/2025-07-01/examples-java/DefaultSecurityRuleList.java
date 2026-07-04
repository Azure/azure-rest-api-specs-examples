
/**
 * Samples for DefaultSecurityRules List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/DefaultSecurityRuleList.json
     */
    /**
     * Sample code: DefaultSecurityRuleList.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void defaultSecurityRuleList(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getDefaultSecurityRules().list("testrg", "nsg1", com.azure.core.util.Context.NONE);
    }
}
