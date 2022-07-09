import com.azure.core.util.Context;

/** Samples for Gateways Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/operationalinsights/resource-manager/Microsoft.OperationalInsights/stable/2020-08-01/examples/WorkspacesGatewaysDelete.json
     */
    /**
     * Sample code: DeleteGateways.
     *
     * @param manager Entry point to LogAnalyticsManager.
     */
    public static void deleteGateways(com.azure.resourcemanager.loganalytics.LogAnalyticsManager manager) {
        manager
            .gateways()
            .deleteWithResponse("OIAutoRest5123", "aztest5048", "00000000-0000-0000-0000-00000000000", Context.NONE);
    }
}
