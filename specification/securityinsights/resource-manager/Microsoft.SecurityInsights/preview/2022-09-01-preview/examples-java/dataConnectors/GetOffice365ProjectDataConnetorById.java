
/**
 * Samples for DataConnectors Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/securityinsights/resource-manager/Microsoft.SecurityInsights/preview/2022-09-01-preview/examples/
     * dataConnectors/GetOffice365ProjectDataConnetorById.json
     */
    /**
     * Sample code: Get an Office365 Project data connector.
     * 
     * @param manager Entry point to SecurityInsightsManager.
     */
    public static void
        getAnOffice365ProjectDataConnector(com.azure.resourcemanager.securityinsights.SecurityInsightsManager manager) {
        manager.dataConnectors().getWithResponse("myRg", "myWorkspace", "73e01a99-5cd7-4139-a149-9f2736ff2ab5",
            com.azure.core.util.Context.NONE);
    }
}
