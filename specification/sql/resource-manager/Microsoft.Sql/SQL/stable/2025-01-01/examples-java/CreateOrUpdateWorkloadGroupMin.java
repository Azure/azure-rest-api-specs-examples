
import com.azure.resourcemanager.sql.fluent.models.WorkloadGroupInner;

/**
 * Samples for WorkloadGroups CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-01-01/CreateOrUpdateWorkloadGroupMin.json
     */
    /**
     * Sample code: Create a workload group with the required properties specified.
     * 
     * @param manager Entry point to SqlServerManager.
     */
    public static void
        createAWorkloadGroupWithTheRequiredPropertiesSpecified(com.azure.resourcemanager.sql.SqlServerManager manager) {
        manager.serviceClient().getWorkloadGroups().createOrUpdate(
            "Default-SQL-SouthEastAsia", "testsvr", "testdb", "smallrc", new WorkloadGroupInner()
                .withMinResourcePercent(0).withMaxResourcePercent(100).withMinResourcePercentPerRequest(3.0),
            com.azure.core.util.Context.NONE);
    }
}
