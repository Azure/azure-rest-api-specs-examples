
import com.azure.core.util.Context;

/** Samples for ActionGroups ListByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/monitor/resource-manager/Microsoft.Insights/stable/2022-06-01/examples/listActionGroups.json
     */
    /**
     * Sample code: List action groups at resource group level.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void listActionGroupsAtResourceGroupLevel(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.diagnosticSettings().manager().serviceClient().getActionGroups()
            .listByResourceGroup("Default-NotificationRules", Context.NONE);
    }
}
