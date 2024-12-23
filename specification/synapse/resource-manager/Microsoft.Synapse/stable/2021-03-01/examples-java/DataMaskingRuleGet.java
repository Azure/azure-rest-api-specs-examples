
/**
 * Samples for DataMaskingRules Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/synapse/resource-manager/Microsoft.Synapse/stable/2021-03-01/examples/DataMaskingRuleGet.json
     */
    /**
     * Sample code: Get data masking rule.
     * 
     * @param manager Entry point to SynapseManager.
     */
    public static void getDataMaskingRule(com.azure.resourcemanager.synapse.SynapseManager manager) {
        manager.dataMaskingRules().getWithResponse("sqlcrudtest-6852", "sqlcrudtest-2080", "sqlcrudtest-331", "rule1",
            com.azure.core.util.Context.NONE);
    }
}
