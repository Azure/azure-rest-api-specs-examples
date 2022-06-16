import com.azure.core.util.Context;

/** Samples for DataConnectors Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/securityinsights/resource-manager/Microsoft.SecurityInsights/preview/2021-09-01-preview/examples/dataConnectors/GetThreatIntelligenceTaxiiById.json
     */
    /**
     * Sample code: Get a TI Taxii data connector.
     *
     * @param manager Entry point to SecurityInsightsManager.
     */
    public static void getATITaxiiDataConnector(
        com.azure.resourcemanager.securityinsights.SecurityInsightsManager manager) {
        manager
            .dataConnectors()
            .getWithResponse("myRg", "myWorkspace", "c39bb458-02a7-4b3f-b0c8-71a1d2692652", Context.NONE);
    }
}
