import com.azure.core.util.Context;

/** Samples for DataConnectors List. */
public final class Main {
    /*
     * x-ms-original-file: specification/securityinsights/resource-manager/Microsoft.SecurityInsights/preview/2021-09-01-preview/examples/dataConnectors/GetDataConnectors.json
     */
    /**
     * Sample code: Get all data connectors.
     *
     * @param manager Entry point to SecurityInsightsManager.
     */
    public static void getAllDataConnectors(
        com.azure.resourcemanager.securityinsights.SecurityInsightsManager manager) {
        manager.dataConnectors().list("myRg", "myWorkspace", Context.NONE);
    }
}
