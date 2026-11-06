
/**
 * Samples for TenantActivityLogs List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2015-04-01/GetTenantActivityLogsSelected.json
     */
    /**
     * Sample code: Get Tenant Activity Logs with select.
     * 
     * @param manager Entry point to MonitorManager.
     */
    public static void getTenantActivityLogsWithSelect(com.azure.resourcemanager.monitor.MonitorManager manager) {
        manager.serviceClient().getTenantActivityLogs().list(null,
            "eventName,id,resourceGroupName,resourceProviderName,operationName,status,eventTimestamp,correlationId,submissionTimestamp,level",
            com.azure.core.util.Context.NONE);
    }
}
