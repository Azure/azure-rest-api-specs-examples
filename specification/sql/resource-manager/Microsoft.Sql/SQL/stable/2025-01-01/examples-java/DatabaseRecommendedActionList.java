
/**
 * Samples for DatabaseRecommendedActions ListByDatabaseAdvisor.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-01-01/DatabaseRecommendedActionList.json
     */
    /**
     * Sample code: List of database recommended actions.
     * 
     * @param manager Entry point to SqlServerManager.
     */
    public static void listOfDatabaseRecommendedActions(com.azure.resourcemanager.sql.SqlServerManager manager) {
        manager.serviceClient().getDatabaseRecommendedActions().listByDatabaseAdvisorWithResponse(
            "workloadinsight-demos", "misosisvr", "IndexAdvisor_test_3", "CreateIndex",
            com.azure.core.util.Context.NONE);
    }
}
