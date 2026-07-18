
import com.azure.resourcemanager.sql.fluent.models.AdvisorInner;
import com.azure.resourcemanager.sql.models.AutoExecuteStatus;

/**
 * Samples for ServerAdvisors Update.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-01-01/ServerAdvisorUpdate.json
     */
    /**
     * Sample code: Update server advisor.
     * 
     * @param manager Entry point to SqlServerManager.
     */
    public static void updateServerAdvisor(com.azure.resourcemanager.sql.SqlServerManager manager) {
        manager.serviceClient().getServerAdvisors().updateWithResponse("workloadinsight-demos", "misosisvr",
            "CreateIndex", new AdvisorInner().withAutoExecuteStatus(AutoExecuteStatus.DISABLED),
            com.azure.core.util.Context.NONE);
    }
}
