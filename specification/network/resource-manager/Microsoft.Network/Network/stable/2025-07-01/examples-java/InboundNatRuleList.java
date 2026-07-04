
/**
 * Samples for InboundNatRules List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/InboundNatRuleList.json
     */
    /**
     * Sample code: InboundNatRuleList.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void inboundNatRuleList(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getInboundNatRules().list("testrg", "lb1", com.azure.core.util.Context.NONE);
    }
}
