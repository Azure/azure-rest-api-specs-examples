
import com.azure.core.util.Context;

/** Samples for ActionGroups Delete. */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/monitor/resource-manager/Microsoft.Insights/stable/2022-06-01/examples/deleteActionGroup.json
     */
    /**
     * Sample code: Delete an action group.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void deleteAnActionGroup(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.diagnosticSettings().manager().serviceClient().getActionGroups()
            .deleteWithResponse("Default-NotificationRules", "SampleActionGroup", Context.NONE);
    }
}
