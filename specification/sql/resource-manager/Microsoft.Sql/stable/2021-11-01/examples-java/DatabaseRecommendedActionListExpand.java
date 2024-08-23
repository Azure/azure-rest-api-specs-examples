
/**
 * Samples for DatabaseAdvisors ListByDatabase.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/sql/resource-manager/Microsoft.Sql/stable/2021-11-01/examples/DatabaseRecommendedActionListExpand.
     * json
     */
    /**
     * Sample code: List of database recommended actions for all advisors.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void
        listOfDatabaseRecommendedActionsForAllAdvisors(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.sqlServers().manager().serviceClient().getDatabaseAdvisors().listByDatabaseWithResponse(
            "workloadinsight-demos", "misosisvr", "IndexAdvisor_test_3", "recommendedActions",
            com.azure.core.util.Context.NONE);
    }
}
