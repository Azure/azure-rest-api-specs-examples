
/**
 * Samples for DataConnectors Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/securityinsights/resource-manager/Microsoft.SecurityInsights/preview/2022-09-01-preview/examples/
     * dataConnectors/GetMicrosoftInsiderRiskManagementById.json
     */
    /**
     * Sample code: Get an Office IRM data connector.
     * 
     * @param manager Entry point to SecurityInsightsManager.
     */
    public static void
        getAnOfficeIRMDataConnector(com.azure.resourcemanager.securityinsights.SecurityInsightsManager manager) {
        manager.dataConnectors().getWithResponse("myRg", "myWorkspace", "3d3e955e-33eb-401d-89a7-251c81ddd660",
            com.azure.core.util.Context.NONE);
    }
}
