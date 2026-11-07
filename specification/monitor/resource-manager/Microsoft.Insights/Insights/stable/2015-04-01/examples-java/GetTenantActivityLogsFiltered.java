
/**
 * Samples for TenantActivityLogs List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2015-04-01/GetTenantActivityLogsFiltered.json
     */
    /**
     * Sample code: Get Tenant Activity Logs with filter.
     * 
     * @param manager Entry point to MonitorManager.
     */
    public static void getTenantActivityLogsWithFilter(com.azure.resourcemanager.monitor.MonitorManager manager) {
        manager.serviceClient().getTenantActivityLogs().list(
            "eventTimestamp ge '2015-01-21T20:00:00Z' and eventTimestamp le '2015-01-23T20:00:00Z' and resourceGroupName eq 'MSSupportGroup'",
            null, com.azure.core.util.Context.NONE);
    }
}
