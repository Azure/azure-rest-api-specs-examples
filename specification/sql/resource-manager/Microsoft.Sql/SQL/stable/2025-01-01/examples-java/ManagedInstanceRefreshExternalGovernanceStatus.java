
/**
 * Samples for ManagedInstances RefreshStatus.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-01-01/ManagedInstanceRefreshExternalGovernanceStatus.json
     */
    /**
     * Sample code: Refresh external governance enablement status.
     * 
     * @param manager Entry point to SqlServerManager.
     */
    public static void
        refreshExternalGovernanceEnablementStatus(com.azure.resourcemanager.sql.SqlServerManager manager) {
        manager.serviceClient().getManagedInstances().refreshStatus("sqlcrudtest-7398", "sqlcrudtest-4645",
            com.azure.core.util.Context.NONE);
    }
}
