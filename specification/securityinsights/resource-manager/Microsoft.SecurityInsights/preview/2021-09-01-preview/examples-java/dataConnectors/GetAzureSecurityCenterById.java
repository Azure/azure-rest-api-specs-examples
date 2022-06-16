import com.azure.core.util.Context;

/** Samples for DataConnectors Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/securityinsights/resource-manager/Microsoft.SecurityInsights/preview/2021-09-01-preview/examples/dataConnectors/GetAzureSecurityCenterById.json
     */
    /**
     * Sample code: Get a ASC data connector.
     *
     * @param manager Entry point to SecurityInsightsManager.
     */
    public static void getAASCDataConnector(
        com.azure.resourcemanager.securityinsights.SecurityInsightsManager manager) {
        manager
            .dataConnectors()
            .getWithResponse("myRg", "myWorkspace", "763f9fa1-c2d3-4fa2-93e9-bccd4899aa12", Context.NONE);
    }
}
