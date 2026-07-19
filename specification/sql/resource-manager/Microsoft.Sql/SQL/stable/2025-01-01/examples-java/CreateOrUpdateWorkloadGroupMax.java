
import com.azure.resourcemanager.sql.fluent.models.WorkloadGroupInner;

/**
 * Samples for WorkloadGroups CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-01-01/CreateOrUpdateWorkloadGroupMax.json
     */
    /**
     * Sample code: Create a workload group with all properties specified.
     * 
     * @param manager Entry point to SqlServerManager.
     */
    public static void
        createAWorkloadGroupWithAllPropertiesSpecified(com.azure.resourcemanager.sql.SqlServerManager manager) {
        manager.serviceClient().getWorkloadGroups().createOrUpdate("Default-SQL-SouthEastAsia", "testsvr", "testdb",
            "smallrc",
            new WorkloadGroupInner().withMinResourcePercent(0).withMaxResourcePercent(100)
                .withMinResourcePercentPerRequest(3.0).withMaxResourcePercentPerRequest(3.0D).withImportance("normal")
                .withQueryExecutionTimeout(0),
            com.azure.core.util.Context.NONE);
    }
}
