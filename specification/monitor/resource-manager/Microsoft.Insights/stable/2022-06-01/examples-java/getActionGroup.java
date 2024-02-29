
/**
 * Samples for ActionGroups GetByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/monitor/resource-manager/Microsoft.Insights/stable/2022-06-01/examples/getActionGroup.json
     */
    /**
     * Sample code: Get an action group.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getAnActionGroup(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.diagnosticSettings().manager().serviceClient().getActionGroups().getByResourceGroupWithResponse(
            "Default-NotificationRules", "SampleActionGroup", com.azure.core.util.Context.NONE);
    }
}
