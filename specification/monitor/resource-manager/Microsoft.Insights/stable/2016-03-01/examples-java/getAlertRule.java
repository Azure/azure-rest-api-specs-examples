
import com.azure.core.util.Context;

/** Samples for AlertRules GetByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/monitor/resource-manager/Microsoft.Insights/stable/2016-03-01/examples/getAlertRule.json
     */
    /**
     * Sample code: Get an alert rule.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getAnAlertRule(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.diagnosticSettings().manager().serviceClient().getAlertRules()
            .getByResourceGroupWithResponse("Rac46PostSwapRG", "chiricutin", Context.NONE);
    }
}
