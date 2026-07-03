
/**
 * Samples for NatRules ListByVpnGateway.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/NatRuleList.json
     */
    /**
     * Sample code: NatRuleList.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void natRuleList(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getNatRules().listByVpnGateway("rg1", "gateway1", com.azure.core.util.Context.NONE);
    }
}
