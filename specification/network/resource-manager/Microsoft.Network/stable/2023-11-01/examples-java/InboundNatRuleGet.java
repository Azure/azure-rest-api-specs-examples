
/**
 * Samples for InboundNatRules Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/network/resource-manager/Microsoft.Network/stable/2023-11-01/examples/InboundNatRuleGet.json
     */
    /**
     * Sample code: InboundNatRuleGet.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void inboundNatRuleGet(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.networks().manager().serviceClient().getInboundNatRules().getWithResponse("testrg", "lb1", "natRule1.1",
            null, com.azure.core.util.Context.NONE);
    }
}
