import com.azure.core.util.Context;

/** Samples for EntityQueries Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/securityinsights/resource-manager/Microsoft.SecurityInsights/preview/2022-01-01-preview/examples/entityQueries/DeleteEntityQuery.json
     */
    /**
     * Sample code: Delete an entity query.
     *
     * @param manager Entry point to SecurityInsightsManager.
     */
    public static void deleteAnEntityQuery(com.azure.resourcemanager.securityinsights.SecurityInsightsManager manager) {
        manager
            .entityQueries()
            .deleteWithResponse("myRg", "myWorkspace", "07da3cc8-c8ad-4710-a44e-334cdcb7882b", Context.NONE);
    }
}
