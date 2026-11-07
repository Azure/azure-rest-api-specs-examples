
/**
 * Samples for TenantActivityLogs List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2015-04-01/GetTenantActivityLogsNoParams.json
     */
    /**
     * Sample code: Get Tenant Activity Logs without filter or select.
     * 
     * @param manager Entry point to MonitorManager.
     */
    public static void
        getTenantActivityLogsWithoutFilterOrSelect(com.azure.resourcemanager.monitor.MonitorManager manager) {
        manager.serviceClient().getTenantActivityLogs().list(null, null, com.azure.core.util.Context.NONE);
    }
}
