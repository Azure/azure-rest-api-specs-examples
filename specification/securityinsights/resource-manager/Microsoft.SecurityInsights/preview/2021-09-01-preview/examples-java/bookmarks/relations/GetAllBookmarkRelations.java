import com.azure.core.util.Context;

/** Samples for BookmarkRelations List. */
public final class Main {
    /*
     * x-ms-original-file: specification/securityinsights/resource-manager/Microsoft.SecurityInsights/preview/2021-09-01-preview/examples/bookmarks/relations/GetAllBookmarkRelations.json
     */
    /**
     * Sample code: Get all bookmark relations.
     *
     * @param manager Entry point to SecurityInsightsManager.
     */
    public static void getAllBookmarkRelations(
        com.azure.resourcemanager.securityinsights.SecurityInsightsManager manager) {
        manager
            .bookmarkRelations()
            .list("myRg", "myWorkspace", "2216d0e1-91e3-4902-89fd-d2df8c535096", null, null, null, null, Context.NONE);
    }
}
