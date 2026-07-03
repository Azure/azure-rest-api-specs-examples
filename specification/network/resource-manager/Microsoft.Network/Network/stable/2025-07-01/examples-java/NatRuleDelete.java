
/**
 * Samples for NatRules Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/NatRuleDelete.json
     */
    /**
     * Sample code: NatRuleDelete.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void natRuleDelete(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getNatRules().delete("rg1", "gateway1", "natRule1", com.azure.core.util.Context.NONE);
    }
}
