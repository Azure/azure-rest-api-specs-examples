
/**
 * Samples for ActivityLogAlerts ListByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/monitor/resource-manager/Microsoft.Insights/stable/2020-10-01/examples/
     * ActivityLogAlertRule_ListByResourceGroupName.json
     */
    /**
     * Sample code: List activity log alerts.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void listActivityLogAlerts(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.diagnosticSettings().manager().serviceClient().getActivityLogAlerts()
            .listByResourceGroup("MyResourceGroup", com.azure.core.util.Context.NONE);
    }
}
