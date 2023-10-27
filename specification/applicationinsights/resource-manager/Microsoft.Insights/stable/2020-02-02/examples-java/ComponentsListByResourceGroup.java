/** Samples for Components ListByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/applicationinsights/resource-manager/Microsoft.Insights/stable/2020-02-02/examples/ComponentsListByResourceGroup.json
     */
    /**
     * Sample code: ComponentListByResourceGroup.
     *
     * @param manager Entry point to ApplicationInsightsManager.
     */
    public static void componentListByResourceGroup(
        com.azure.resourcemanager.applicationinsights.ApplicationInsightsManager manager) {
        manager.components().listByResourceGroup("my-resource-group", com.azure.core.util.Context.NONE);
    }
}
