
/**
 * Samples for DiscoverySolution List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2024-03-01-preview/ListDiscoverySolutionsAtTenantScope.json
     */
    /**
     * Sample code: List DiscoverySolutions at resource scope.
     * 
     * @param manager Entry point to SelfHelpManager.
     */
    public static void
        listDiscoverySolutionsAtResourceScope(com.azure.resourcemanager.selfhelp.SelfHelpManager manager) {
        manager.discoverySolutions().list("ProblemClassificationId eq 'SampleProblemClassificationId1'", null,
            com.azure.core.util.Context.NONE);
    }
}
