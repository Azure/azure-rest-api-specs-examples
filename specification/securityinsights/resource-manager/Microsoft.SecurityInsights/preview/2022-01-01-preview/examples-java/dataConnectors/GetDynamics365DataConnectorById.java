import com.azure.core.util.Context;

/** Samples for DataConnectors Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/securityinsights/resource-manager/Microsoft.SecurityInsights/preview/2022-01-01-preview/examples/dataConnectors/GetDynamics365DataConnectorById.json
     */
    /**
     * Sample code: Get a Dynamics365 data connector.
     *
     * @param manager Entry point to SecurityInsightsManager.
     */
    public static void getADynamics365DataConnector(
        com.azure.resourcemanager.securityinsights.SecurityInsightsManager manager) {
        manager
            .dataConnectors()
            .getWithResponse("myRg", "myWorkspace", "c2541efb-c9a6-47fe-9501-87d1017d1512", Context.NONE);
    }
}
