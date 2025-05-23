
/**
 * Samples for DataConnectors List.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/securityinsights/resource-manager/Microsoft.SecurityInsights/stable/2022-11-01/examples/
     * dataConnectors/GetDataConnectors.json
     */
    /**
     * Sample code: Get all data connectors.
     * 
     * @param manager Entry point to SecurityInsightsManager.
     */
    public static void
        getAllDataConnectors(com.azure.resourcemanager.securityinsights.SecurityInsightsManager manager) {
        manager.dataConnectors().list("myRg", "myWorkspace", com.azure.core.util.Context.NONE);
    }
}
