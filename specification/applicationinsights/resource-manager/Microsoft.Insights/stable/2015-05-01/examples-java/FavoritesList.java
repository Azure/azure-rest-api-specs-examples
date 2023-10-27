/** Samples for Favorites List. */
public final class Main {
    /*
     * x-ms-original-file: specification/applicationinsights/resource-manager/Microsoft.Insights/stable/2015-05-01/examples/FavoritesList.json
     */
    /**
     * Sample code: FavoritesList.
     *
     * @param manager Entry point to ApplicationInsightsManager.
     */
    public static void favoritesList(com.azure.resourcemanager.applicationinsights.ApplicationInsightsManager manager) {
        manager
            .favorites()
            .listWithResponse(
                "my-resource-group", "my-ai-component", null, null, null, null, com.azure.core.util.Context.NONE);
    }
}
