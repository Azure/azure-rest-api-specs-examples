/** Samples for DataMaskingRules ListBySqlPool. */
public final class Main {
    /*
     * x-ms-original-file: specification/synapse/resource-manager/Microsoft.Synapse/stable/2021-06-01/examples/DataMaskingRuleList.json
     */
    /**
     * Sample code: List data masking rules.
     *
     * @param manager Entry point to SynapseManager.
     */
    public static void listDataMaskingRules(com.azure.resourcemanager.synapse.SynapseManager manager) {
        manager
            .dataMaskingRules()
            .listBySqlPool("sqlcrudtest-6852", "sqlcrudtest-2080", "sqlcrudtest-331", com.azure.core.util.Context.NONE);
    }
}
