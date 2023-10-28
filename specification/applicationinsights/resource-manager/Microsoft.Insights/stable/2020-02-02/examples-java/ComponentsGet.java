/** Samples for Components GetByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/applicationinsights/resource-manager/Microsoft.Insights/stable/2020-02-02/examples/ComponentsGet.json
     */
    /**
     * Sample code: ComponentGet.
     *
     * @param manager Entry point to ApplicationInsightsManager.
     */
    public static void componentGet(com.azure.resourcemanager.applicationinsights.ApplicationInsightsManager manager) {
        manager
            .components()
            .getByResourceGroupWithResponse("my-resource-group", "my-component", com.azure.core.util.Context.NONE);
    }
}
