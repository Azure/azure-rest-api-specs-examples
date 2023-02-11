/** Samples for DataMaskingPolicies Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/synapse/resource-manager/Microsoft.Synapse/stable/2021-06-01/examples/DataMaskingPolicyGet.json
     */
    /**
     * Sample code: Get data masking policy.
     *
     * @param manager Entry point to SynapseManager.
     */
    public static void getDataMaskingPolicy(com.azure.resourcemanager.synapse.SynapseManager manager) {
        manager
            .dataMaskingPolicies()
            .getWithResponse(
                "sqlcrudtest-6852", "sqlcrudtest-2080", "sqlcrudtest-331", com.azure.core.util.Context.NONE);
    }
}
