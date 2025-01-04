
/**
 * Samples for Bookmarks List.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/securityinsights/resource-manager/Microsoft.SecurityInsights/stable/2022-11-01/examples/bookmarks/
     * GetBookmarks.json
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
