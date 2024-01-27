
import com.azure.core.util.Context;

/** Samples for AutoscaleSettings Delete. */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/monitor/resource-manager/Microsoft.Insights/stable/2022-10-01/examples/deleteAutoscaleSetting.json
     */
    /**
     * Sample code: Delete an autoscale setting.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void deleteAnAutoscaleSetting(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.diagnosticSettings().manager().serviceClient().getAutoscaleSettings()
            .deleteWithResponse("TestingMetricsScaleSet", "MySetting", Context.NONE);
    }
}
