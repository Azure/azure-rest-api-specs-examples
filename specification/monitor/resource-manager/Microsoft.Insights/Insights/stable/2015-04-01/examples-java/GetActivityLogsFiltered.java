
/**
 * Samples for ActivityLogs List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2015-04-01/GetActivityLogsFiltered.json
     */
    /**
     * Sample code: Get Activity Logs with filter.
     * 
     * @param manager Entry point to MonitorManager.
     */
    public static void getActivityLogsWithFilter(com.azure.resourcemanager.monitor.MonitorManager manager) {
        manager.serviceClient().getActivityLogs().list(
            "eventTimestamp ge '2015-01-21T20:00:00Z' and eventTimestamp le '2015-01-23T20:00:00Z' and resourceGroupName eq 'MSSupportGroup'",
            null, com.azure.core.util.Context.NONE);
    }
}
