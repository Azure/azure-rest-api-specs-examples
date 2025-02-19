
/**
 * Samples for SqlPoolWorkloadGroup CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/synapse/resource-manager/Microsoft.Synapse/stable/2021-03-01/examples/
     * CreateOrUpdateSqlPoolWorkloadGroupMin.json
     */
    /**
     * Sample code: Create a workload group with the required properties specified.
     * 
     * @param manager Entry point to SynapseManager.
     */
    public static void createAWorkloadGroupWithTheRequiredPropertiesSpecified(
        com.azure.resourcemanager.synapse.SynapseManager manager) {
        manager.sqlPoolWorkloadGroups().define("smallrc")
            .withExistingSqlPool("sqlcrudtest-6852", "sqlcrudtest-2080", "sqlcrudtest-9187").withMinResourcePercent(0)
            .withMaxResourcePercent(100).withMinResourcePercentPerRequest(3.0).create();
    }
}
