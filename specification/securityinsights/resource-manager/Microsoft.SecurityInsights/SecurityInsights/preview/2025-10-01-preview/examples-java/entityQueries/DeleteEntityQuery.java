
/**
 * Samples for EntityQueries Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-10-01-preview/entityQueries/DeleteEntityQuery.json
     */
    /**
     * Sample code: Delete an entity query.
     * 
     * @param manager Entry point to SecurityInsightsManager.
     */
    public static void deleteAnEntityQuery(com.azure.resourcemanager.securityinsights.SecurityInsightsManager manager) {
        manager.entityQueries().deleteWithResponse("myRg", "myWorkspace", "07da3cc8-c8ad-4710-a44e-334cdcb7882b",
            com.azure.core.util.Context.NONE);
    }
}
