
/**
 * Samples for WorkloadGroups ListByDatabase.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/sql/resource-manager/Microsoft.Sql/stable/2021-11-01/examples/GetWorkloadGroupList.json
     */
    /**
     * Sample code: Get the list of workload groups for a data warehouse.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void
        getTheListOfWorkloadGroupsForADataWarehouse(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.sqlServers().manager().serviceClient().getWorkloadGroups().listByDatabase("Default-SQL-SouthEastAsia",
            "testsvr", "testdb", com.azure.core.util.Context.NONE);
    }
}
