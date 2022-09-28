import com.azure.core.util.Context;

/** Samples for DataConnectors Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/securityinsights/resource-manager/Microsoft.SecurityInsights/preview/2022-09-01-preview/examples/dataConnectors/GetIoTById.json
     */
    /**
     * Sample code: Get a IoT data connector.
     *
     * @param manager Entry point to SecurityInsightsManager.
     */
    public static void getAIoTDataConnector(
        com.azure.resourcemanager.securityinsights.SecurityInsightsManager manager) {
        manager
            .dataConnectors()
            .getWithResponse("myRg", "myWorkspace", "d2e5dc7a-f3a2-429d-954b-939fa8c2932e", Context.NONE);
    }
}
