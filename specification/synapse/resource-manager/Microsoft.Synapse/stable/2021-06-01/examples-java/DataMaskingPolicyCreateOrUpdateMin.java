import com.azure.resourcemanager.synapse.models.DataMaskingState;

/** Samples for DataMaskingPolicies CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/synapse/resource-manager/Microsoft.Synapse/stable/2021-06-01/examples/DataMaskingPolicyCreateOrUpdateMin.json
     */
    /**
     * Sample code: Create or update data masking policy min.
     *
     * @param manager Entry point to SynapseManager.
     */
    public static void createOrUpdateDataMaskingPolicyMin(com.azure.resourcemanager.synapse.SynapseManager manager) {
        manager
            .dataMaskingPolicies()
            .define()
            .withExistingSqlPool("sqlcrudtest-6852", "sqlcrudtest-2080", "sqlcrudtest-331")
            .withDataMaskingState(DataMaskingState.ENABLED)
            .create();
    }
}
