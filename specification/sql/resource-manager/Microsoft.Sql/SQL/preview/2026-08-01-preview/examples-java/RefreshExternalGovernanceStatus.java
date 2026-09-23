
/**
 * Samples for Servers RefreshStatus.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-08-01-preview/RefreshExternalGovernanceStatus.json
     */
    /**
     * Sample code: Refresh external governance enablement status.
     * 
     * @param manager Entry point to SqlServerManager.
     */
    public static void
        refreshExternalGovernanceEnablementStatus(com.azure.resourcemanager.sql.SqlServerManager manager) {
        manager.serviceClient().getServers().refreshStatus("sqlcrudtest-7398", "sqlcrudtest-4645",
            com.azure.core.util.Context.NONE);
    }
}
