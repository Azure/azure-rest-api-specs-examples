
import com.azure.core.util.Context;

/** Samples for AutoscaleSettings GetByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/monitor/resource-manager/Microsoft.Insights/stable/2022-10-01/examples/getAutoscaleSetting.json
     */
    /**
     * Sample code: Get an autoscale setting.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getAnAutoscaleSetting(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.diagnosticSettings().manager().serviceClient().getAutoscaleSettings()
            .getByResourceGroupWithResponse("TestingMetricsScaleSet", "MySetting", Context.NONE);
    }
}
