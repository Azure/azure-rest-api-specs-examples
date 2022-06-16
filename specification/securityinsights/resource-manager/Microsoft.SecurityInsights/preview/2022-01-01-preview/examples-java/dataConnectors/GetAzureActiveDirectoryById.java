import com.azure.core.util.Context;

/** Samples for DataConnectors Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/securityinsights/resource-manager/Microsoft.SecurityInsights/preview/2022-01-01-preview/examples/dataConnectors/GetAzureActiveDirectoryById.json
     */
    /**
     * Sample code: Get an AAD data connector.
     *
     * @param manager Entry point to SecurityInsightsManager.
     */
    public static void getAnAADDataConnector(
        com.azure.resourcemanager.securityinsights.SecurityInsightsManager manager) {
        manager
            .dataConnectors()
            .getWithResponse("myRg", "myWorkspace", "f0cd27d2-5f03-4c06-ba31-d2dc82dcb51d", Context.NONE);
    }
}
