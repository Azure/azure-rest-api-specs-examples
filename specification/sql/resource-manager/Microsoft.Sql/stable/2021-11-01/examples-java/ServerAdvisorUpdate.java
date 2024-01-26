
import com.azure.core.util.Context;
import com.azure.resourcemanager.sql.fluent.models.AdvisorInner;
import com.azure.resourcemanager.sql.models.AutoExecuteStatus;

/** Samples for ServerAdvisors Update. */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/sql/resource-manager/Microsoft.Sql/stable/2021-11-01/examples/ServerAdvisorUpdate.json
     */
    /**
     * Sample code: Update server advisor.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void updateServerAdvisor(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.sqlServers().manager().serviceClient().getServerAdvisors().updateWithResponse("workloadinsight-demos",
            "misosisvr", "CreateIndex", new AdvisorInner().withAutoExecuteStatus(AutoExecuteStatus.DISABLED),
            Context.NONE);
    }
}
