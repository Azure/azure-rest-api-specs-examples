
/**
 * Samples for BookmarkRelations List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-10-01-preview/bookmarks/relations/GetAllBookmarkRelations.json
     */
    /**
     * Sample code: Get all bookmark relations.
     * 
     * @param manager Entry point to SecurityInsightsManager.
     */
    public static void
        getAllBookmarkRelations(com.azure.resourcemanager.securityinsights.SecurityInsightsManager manager) {
        manager.bookmarkRelations().list("myRg", "myWorkspace", "2216d0e1-91e3-4902-89fd-d2df8c535096", null, null,
            null, null, com.azure.core.util.Context.NONE);
    }
}
