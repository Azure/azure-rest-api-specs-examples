import com.azure.core.util.Context;

/** Samples for EntityQueries Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/securityinsights/resource-manager/Microsoft.SecurityInsights/preview/2022-09-01-preview/examples/entityQueries/GetExpansionEntityQueryById.json
     */
    /**
     * Sample code: Get an Expansion entity query.
     *
     * @param manager Entry point to SecurityInsightsManager.
     */
    public static void getAnExpansionEntityQuery(
        com.azure.resourcemanager.securityinsights.SecurityInsightsManager manager) {
        manager
            .entityQueries()
            .getWithResponse("myRg", "myWorkspace", "07da3cc8-c8ad-4710-a44e-334cdcb7882b", Context.NONE);
    }
}
