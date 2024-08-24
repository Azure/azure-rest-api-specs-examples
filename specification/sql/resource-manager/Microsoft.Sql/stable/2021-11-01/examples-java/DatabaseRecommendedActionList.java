
/**
 * Samples for DatabaseRecommendedActions ListByDatabaseAdvisor.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/sql/resource-manager/Microsoft.Sql/stable/2021-11-01/examples/DatabaseRecommendedActionList.json
     */
    /**
     * Sample code: List of database recommended actions.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void listOfDatabaseRecommendedActions(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.sqlServers().manager().serviceClient().getDatabaseRecommendedActions().listByDatabaseAdvisorWithResponse(
            "workloadinsight-demos", "misosisvr", "IndexAdvisor_test_3", "CreateIndex",
            com.azure.core.util.Context.NONE);
    }
}
