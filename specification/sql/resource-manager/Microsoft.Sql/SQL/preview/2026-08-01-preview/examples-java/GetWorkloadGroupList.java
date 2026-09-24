
/**
 * Samples for WorkloadGroups ListByDatabase.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-08-01-preview/GetWorkloadGroupList.json
     */
    /**
     * Sample code: Get the list of workload groups for a data warehouse.
     * 
     * @param manager Entry point to SqlServerManager.
     */
    public static void
        getTheListOfWorkloadGroupsForADataWarehouse(com.azure.resourcemanager.sql.SqlServerManager manager) {
        manager.serviceClient().getWorkloadGroups().listByDatabase("Default-SQL-SouthEastAsia", "testsvr", "testdb",
            com.azure.core.util.Context.NONE);
    }
}
