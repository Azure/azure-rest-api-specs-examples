
import com.azure.core.util.Context;

/** Samples for ActivityLogs List. */
public final class Main {
    /*
     * x-ms-original-file: specification/monitor/resource-manager/Microsoft.Insights/stable/2015-04-01/examples/
     * GetActivityLogsFilteredAndSelected.json
     */
    /**
     * Sample code: Get Activity Logs with filter and select.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getActivityLogsWithFilterAndSelect(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.diagnosticSettings().manager().serviceClient().getActivityLogs().list(
            "eventTimestamp ge '2015-01-21T20:00:00Z' and eventTimestamp le '2015-01-23T20:00:00Z' and"
                + " resourceGroupName eq 'MSSupportGroup'",
            "eventName,id,resourceGroupName,resourceProviderName,operationName,status,eventTimestamp,correlationId,submissionTimestamp,level",
            Context.NONE);
    }
}
