
/**
 * Samples for NatRules Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/NatRuleGet.json
     */
    /**
     * Sample code: NatRuleGet.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void natRuleGet(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getNatRules().getWithResponse("rg1", "gateway1", "natRule1",
            com.azure.core.util.Context.NONE);
    }
}
