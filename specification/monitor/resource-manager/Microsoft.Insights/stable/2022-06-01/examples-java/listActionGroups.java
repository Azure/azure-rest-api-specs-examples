
/**
 * Samples for ActionGroups List.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/monitor/resource-manager/Microsoft.Insights/stable/2022-06-01/examples/listActionGroups.json
     */
    /**
     * Sample code: List action groups at subscription level.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void listActionGroupsAtSubscriptionLevel(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.diagnosticSettings().manager().serviceClient().getActionGroups().list(com.azure.core.util.Context.NONE);
    }
}
