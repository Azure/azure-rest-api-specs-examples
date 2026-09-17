
/**
 * Samples for Bookmarks List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-10-01-preview/bookmarks/GetBookmarks.json
     */
    /**
     * Sample code: Get all bookmarks.
     * 
     * @param manager Entry point to SecurityInsightsManager.
     */
    public static void getAllBookmarks(com.azure.resourcemanager.securityinsights.SecurityInsightsManager manager) {
        manager.bookmarks().list("myRg", "myWorkspace", com.azure.core.util.Context.NONE);
    }
}
