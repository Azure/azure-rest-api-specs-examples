import com.azure.core.util.Context;

/** Samples for DataConnectors Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/securityinsights/resource-manager/Microsoft.SecurityInsights/preview/2022-01-01-preview/examples/dataConnectors/GetMicrosoftCloudAppSecurityById.json
     */
    /**
     * Sample code: Get a MCAS data connector.
     *
     * @param manager Entry point to SecurityInsightsManager.
     */
    public static void getAMCASDataConnector(
        com.azure.resourcemanager.securityinsights.SecurityInsightsManager manager) {
        manager
            .dataConnectors()
            .getWithResponse("myRg", "myWorkspace", "b96d014d-b5c2-4a01-9aba-a8058f629d42", Context.NONE);
    }
}
