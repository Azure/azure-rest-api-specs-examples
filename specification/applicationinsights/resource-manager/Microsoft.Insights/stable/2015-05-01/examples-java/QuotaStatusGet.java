/** Samples for ComponentQuotaStatus Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/applicationinsights/resource-manager/Microsoft.Insights/stable/2015-05-01/examples/QuotaStatusGet.json
     */
    /**
     * Sample code: ComponentCurrentBillingFeaturesGet.
     *
     * @param manager Entry point to ApplicationInsightsManager.
     */
    public static void componentCurrentBillingFeaturesGet(
        com.azure.resourcemanager.applicationinsights.ApplicationInsightsManager manager) {
        manager
            .componentQuotaStatus()
            .getWithResponse("my-resource-group", "my-component", com.azure.core.util.Context.NONE);
    }
}
