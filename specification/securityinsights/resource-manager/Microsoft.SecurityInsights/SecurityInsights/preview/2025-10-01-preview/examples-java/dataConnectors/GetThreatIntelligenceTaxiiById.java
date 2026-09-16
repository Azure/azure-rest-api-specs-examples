
/**
 * Samples for DataConnectors Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-10-01-preview/dataConnectors/GetThreatIntelligenceTaxiiById.json
     */
    /**
     * Sample code: Get a TI Taxii data connector.
     * 
     * @param manager Entry point to SecurityInsightsManager.
     */
    public static void
        getATITaxiiDataConnector(com.azure.resourcemanager.securityinsights.SecurityInsightsManager manager) {
        manager.dataConnectors().getWithResponse("myRg", "myWorkspace", "c39bb458-02a7-4b3f-b0c8-71a1d2692652",
            com.azure.core.util.Context.NONE);
    }
}
