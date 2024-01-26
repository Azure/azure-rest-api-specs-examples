
import com.azure.core.util.Context;

/** Samples for AlertRules Delete. */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/monitor/resource-manager/Microsoft.Insights/stable/2016-03-01/examples/deleteAlertRule.json
     */
    /**
     * Sample code: Delete an alert rule.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void deleteAnAlertRule(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.diagnosticSettings().manager().serviceClient().getAlertRules().deleteWithResponse("Rac46PostSwapRG",
            "chiricutin", Context.NONE);
    }
}
