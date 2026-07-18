
/**
 * Samples for WorkloadGroups Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-01-01/DeleteWorkloadGroup.json
     */
    /**
     * Sample code: Delete a workload group.
     * 
     * @param manager Entry point to SqlServerManager.
     */
    public static void deleteAWorkloadGroup(com.azure.resourcemanager.sql.SqlServerManager manager) {
        manager.serviceClient().getWorkloadGroups().delete("Default-SQL-SouthEastAsia", "testsvr", "testdb",
            "wlm_workloadgroup", com.azure.core.util.Context.NONE);
    }
}
