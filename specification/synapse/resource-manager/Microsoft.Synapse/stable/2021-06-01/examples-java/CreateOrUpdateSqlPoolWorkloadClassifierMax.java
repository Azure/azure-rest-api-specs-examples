/** Samples for SqlPoolWorkloadClassifier CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/synapse/resource-manager/Microsoft.Synapse/stable/2021-06-01/examples/CreateOrUpdateSqlPoolWorkloadClassifierMax.json
     */
    /**
     * Sample code: Create a workload classifier with all properties specified.
     *
     * @param manager Entry point to SynapseManager.
     */
    public static void createAWorkloadClassifierWithAllPropertiesSpecified(
        com.azure.resourcemanager.synapse.SynapseManager manager) {
        manager
            .sqlPoolWorkloadClassifiers()
            .define("wlm_workloadclassifier")
            .withExistingWorkloadGroup("sqlcrudtest-6852", "sqlcrudtest-2080", "sqlcrudtest-9187", "wlm_workloadgroup")
            .withMemberName("dbo")
            .withLabel("test_label")
            .withContext("test_context")
            .withStartTime("12:00")
            .withEndTime("14:00")
            .withImportance("high")
            .create();
    }
}
