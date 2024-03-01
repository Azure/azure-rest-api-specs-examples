
/**
 * Samples for TenantActivityLogs List.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/monitor/resource-manager/Microsoft.Insights/stable/2015-04-01/examples/
     * GetTenantActivityLogsSelected.json
     */
    /**
     * Sample code: Get Tenant Activity Logs with select.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getTenantActivityLogsWithSelect(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.diagnosticSettings().manager().serviceClient().getTenantActivityLogs().list(null,
            "eventName,id,resourceGroupName,resourceProviderName,operationName,status,eventTimestamp,correlationId,submissionTimestamp,level",
            com.azure.core.util.Context.NONE);
    }

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
        azure.diagnosticSettings().manager().serviceClient().getTenantActivityLogs().list(null, null,
            com.azure.core.util.Context.NONE);
    }
}
