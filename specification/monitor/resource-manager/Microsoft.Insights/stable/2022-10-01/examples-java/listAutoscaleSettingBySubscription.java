
/**
 * Samples for AutoscaleSettings List.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/monitor/resource-manager/Microsoft.Insights/stable/2022-10-01/examples/
     * listAutoscaleSettingBySubscription.json
     */
    /**
     * Sample code: List autoscale settings by subs.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void listAutoscaleSettingsBySubs(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.diagnosticSettings().manager().serviceClient().getAutoscaleSettings()
            .list(com.azure.core.util.Context.NONE);
    }
}
