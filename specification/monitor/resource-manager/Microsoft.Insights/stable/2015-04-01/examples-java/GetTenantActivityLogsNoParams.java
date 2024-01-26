
import com.azure.core.util.Context;

/** Samples for TenantActivityLogs List. */
public final class Main {
    /*
     * x-ms-original-file: specification/monitor/resource-manager/Microsoft.Insights/stable/2015-04-01/examples/
     * GetTenantActivityLogsNoParams.json
     */
    /**
     * Sample code: Get Tenant Activity Logs without filter or select.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void
        getTenantActivityLogsWithoutFilterOrSelect(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.diagnosticSettings().manager().serviceClient().getTenantActivityLogs().list(null, null, Context.NONE);
    }
}
