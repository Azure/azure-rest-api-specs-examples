
import com.azure.core.util.Context;

/** Samples for AutoscaleSettings ListByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/monitor/resource-manager/Microsoft.Insights/stable/2022-10-01/examples/listAutoscaleSetting.json
     */
    /**
     * Sample code: List autoscale settings.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void listAutoscaleSettings(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.diagnosticSettings().manager().serviceClient().getAutoscaleSettings()
            .listByResourceGroup("TestingMetricsScaleSet", Context.NONE);
    }
}
