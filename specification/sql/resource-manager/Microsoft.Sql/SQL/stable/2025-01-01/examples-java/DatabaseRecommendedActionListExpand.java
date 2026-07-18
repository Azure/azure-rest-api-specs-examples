
/**
 * Samples for DatabaseAdvisors ListByDatabase.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-01-01/DatabaseRecommendedActionListExpand.json
     */
    /**
     * Sample code: List of database recommended actions for all advisors.
     * 
     * @param manager Entry point to SqlServerManager.
     */
    public static void
        listOfDatabaseRecommendedActionsForAllAdvisors(com.azure.resourcemanager.sql.SqlServerManager manager) {
        manager.serviceClient().getDatabaseAdvisors().listByDatabaseWithResponse("workloadinsight-demos", "misosisvr",
            "IndexAdvisor_test_3", "recommendedActions", com.azure.core.util.Context.NONE);
    }
}
