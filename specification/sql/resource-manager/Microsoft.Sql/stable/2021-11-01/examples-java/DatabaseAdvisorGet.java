
/**
 * Samples for DatabaseAdvisors Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/sql/resource-manager/Microsoft.Sql/stable/2021-11-01/examples/DatabaseAdvisorGet.json
     */
    /**
     * Sample code: Get database advisor.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getDatabaseAdvisor(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.sqlServers().manager().serviceClient().getDatabaseAdvisors().getWithResponse("workloadinsight-demos",
            "misosisvr", "IndexAdvisor_test_3", "CreateIndex", com.azure.core.util.Context.NONE);
    }
}
