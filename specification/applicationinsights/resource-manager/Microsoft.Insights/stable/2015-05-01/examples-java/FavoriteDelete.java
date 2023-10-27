/** Samples for Favorites Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/applicationinsights/resource-manager/Microsoft.Insights/stable/2015-05-01/examples/FavoriteDelete.json
     */
    /**
     * Sample code: FavoriteList.
     *
     * @param manager Entry point to ApplicationInsightsManager.
     */
    public static void favoriteList(com.azure.resourcemanager.applicationinsights.ApplicationInsightsManager manager) {
        manager
            .favorites()
            .deleteWithResponse(
                "my-resource-group",
                "my-ai-component",
                "deadb33f-5e0d-4064-8ebb-1a4ed0313eb2",
                com.azure.core.util.Context.NONE);
    }
}
