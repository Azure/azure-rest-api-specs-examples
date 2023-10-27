/** Samples for Components Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/applicationinsights/resource-manager/Microsoft.Insights/stable/2020-02-02/examples/ComponentsDelete.json
     */
    /**
     * Sample code: ComponentsDelete.
     *
     * @param manager Entry point to ApplicationInsightsManager.
     */
    public static void componentsDelete(
        com.azure.resourcemanager.applicationinsights.ApplicationInsightsManager manager) {
        manager
            .components()
            .deleteByResourceGroupWithResponse("my-resource-group", "my-component", com.azure.core.util.Context.NONE);
    }
}
