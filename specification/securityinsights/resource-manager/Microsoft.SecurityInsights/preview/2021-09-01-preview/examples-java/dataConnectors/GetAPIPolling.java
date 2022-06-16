import com.azure.core.util.Context;

/** Samples for DataConnectors Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/securityinsights/resource-manager/Microsoft.SecurityInsights/preview/2021-09-01-preview/examples/dataConnectors/GetAPIPolling.json
     */
    /**
     * Sample code: Get a APIPolling data connector.
     *
     * @param manager Entry point to SecurityInsightsManager.
     */
    public static void getAAPIPollingDataConnector(
        com.azure.resourcemanager.securityinsights.SecurityInsightsManager manager) {
        manager
            .dataConnectors()
            .getWithResponse("myRg", "myWorkspace", "316ec55e-7138-4d63-ab18-90c8a60fd1c8", Context.NONE);
    }
}
