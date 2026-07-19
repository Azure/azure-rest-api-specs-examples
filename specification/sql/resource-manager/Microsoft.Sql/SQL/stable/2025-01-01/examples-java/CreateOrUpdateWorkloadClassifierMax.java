
import com.azure.resourcemanager.sql.fluent.models.WorkloadClassifierInner;

/**
 * Samples for WorkloadClassifiers CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-01-01/CreateOrUpdateWorkloadClassifierMax.json
     */
    /**
     * Sample code: Create a workload group with all properties specified.
     * 
     * @param manager Entry point to SqlServerManager.
     */
    public static void
        createAWorkloadGroupWithAllPropertiesSpecified(com.azure.resourcemanager.sql.SqlServerManager manager) {
        manager.serviceClient().getWorkloadClassifiers()
            .createOrUpdate("Default-SQL-SouthEastAsia", "testsvr", "testdb", "wlm_workloadgroup",
                "wlm_workloadclassifier",
                new WorkloadClassifierInner().withMemberName("dbo").withLabel("test_label").withContext("test_context")
                    .withStartTime("12:00").withEndTime("14:00").withImportance("high"),
                com.azure.core.util.Context.NONE);
    }
}
