
import com.azure.resourcemanager.sql.fluent.models.WorkloadClassifierInner;

/**
 * Samples for WorkloadClassifiers CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-08-01-preview/CreateOrUpdateWorkloadClassifierMin.json
     */
    /**
     * Sample code: Create a workload group with the required properties specified.
     * 
     * @param manager Entry point to SqlServerManager.
     */
    public static void
        createAWorkloadGroupWithTheRequiredPropertiesSpecified(com.azure.resourcemanager.sql.SqlServerManager manager) {
        manager.serviceClient().getWorkloadClassifiers().createOrUpdate("Default-SQL-SouthEastAsia", "testsvr",
            "testdb", "wlm_workloadgroup", "wlm_workloadclassifier",
            new WorkloadClassifierInner().withMemberName("dbo"), com.azure.core.util.Context.NONE);
    }
}
