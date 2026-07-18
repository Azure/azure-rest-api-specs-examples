
/**
 * Samples for DatabaseAdvisors ListByDatabase.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-01-01/DatabaseAdvisorList.json
     */
    /**
     * Sample code: List of database advisors.
     * 
     * @param manager Entry point to SqlServerManager.
     */
    public static void listOfDatabaseAdvisors(com.azure.resourcemanager.sql.SqlServerManager manager) {
        manager.serviceClient().getDatabaseAdvisors().listByDatabaseWithResponse("workloadinsight-demos", "misosisvr",
            "IndexAdvisor_test_3", null, com.azure.core.util.Context.NONE);
    }
}
