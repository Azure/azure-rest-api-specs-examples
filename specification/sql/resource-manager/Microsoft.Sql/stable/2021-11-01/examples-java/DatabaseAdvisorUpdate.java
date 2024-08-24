
import com.azure.resourcemanager.sql.fluent.models.AdvisorInner;
import com.azure.resourcemanager.sql.models.AutoExecuteStatus;

/**
 * Samples for DatabaseAdvisors Update.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/sql/resource-manager/Microsoft.Sql/stable/2021-11-01/examples/DatabaseAdvisorUpdate.json
     */
    /**
     * Sample code: Update database advisor.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void updateDatabaseAdvisor(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.sqlServers().manager().serviceClient().getDatabaseAdvisors().updateWithResponse("workloadinsight-demos",
            "misosisvr", "IndexAdvisor_test_3", "CreateIndex",
            new AdvisorInner().withAutoExecuteStatus(AutoExecuteStatus.DISABLED), com.azure.core.util.Context.NONE);
    }
}
