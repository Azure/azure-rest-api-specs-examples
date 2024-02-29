
/**
 * Samples for ActivityLogAlerts Delete.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/monitor/resource-manager/Microsoft.Insights/stable/2020-10-01/examples/ActivityLogAlertRule_Delete.
     * json
     */
    /**
     * Sample code: Delete an Activity Log Alert rule.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void deleteAnActivityLogAlertRule(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.diagnosticSettings().manager().serviceClient().getActivityLogAlerts()
            .deleteWithResponse("MyResourceGroup", "SampleActivityLogAlertRule", com.azure.core.util.Context.NONE);
    }
}
