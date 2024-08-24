
/**
 * Samples for WorkloadClassifiers Delete.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/sql/resource-manager/Microsoft.Sql/stable/2021-11-01/examples/DeleteWorkloadClassifier.json
     */
    /**
     * Sample code: Delete a workload classifier.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void deleteAWorkloadClassifier(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.sqlServers().manager().serviceClient().getWorkloadClassifiers().delete("Default-SQL-SouthEastAsia",
            "testsvr", "testdb", "wlm_workloadgroup", "wlm_workloadclassifier", com.azure.core.util.Context.NONE);
    }
}
