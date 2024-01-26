
import com.azure.core.util.Context;

/** Samples for ScheduledQueryRules List. */
public final class Main {
    /*
     * x-ms-original-file: specification/monitor/resource-manager/Microsoft.Insights/stable/2022-06-15/examples/
     * listScheduledQueryRulesBySubscription.json
     */
    /**
     * Sample code: List scheduled query rules by subscription.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void listScheduledQueryRulesBySubscription(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.diagnosticSettings().manager().serviceClient().getScheduledQueryRules().list(Context.NONE);
    }
}
