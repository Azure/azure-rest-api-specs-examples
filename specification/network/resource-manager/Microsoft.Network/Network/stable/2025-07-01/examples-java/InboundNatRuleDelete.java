
/**
 * Samples for InboundNatRules Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/InboundNatRuleDelete.json
     */
    /**
     * Sample code: InboundNatRuleDelete.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void inboundNatRuleDelete(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getInboundNatRules().delete("testrg", "lb1", "natRule1.1",
            com.azure.core.util.Context.NONE);
    }
}
