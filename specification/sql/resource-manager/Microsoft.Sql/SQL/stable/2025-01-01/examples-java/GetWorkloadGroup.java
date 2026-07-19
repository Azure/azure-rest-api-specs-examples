
/**
 * Samples for WorkloadGroups Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-01-01/GetWorkloadGroup.json
     */
    /**
     * Sample code: Gets a workload group for a data warehouse.
     * 
     * @param manager Entry point to SqlServerManager.
     */
    public static void getsAWorkloadGroupForADataWarehouse(com.azure.resourcemanager.sql.SqlServerManager manager) {
        manager.serviceClient().getWorkloadGroups().getWithResponse("Default-SQL-SouthEastAsia", "testsvr", "testdb",
            "smallrc", com.azure.core.util.Context.NONE);
    }
}
