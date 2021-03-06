import com.azure.core.util.Context;

/** Samples for LiveToken Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/applicationinsights/resource-manager/Microsoft.Insights/preview/2020-06-02-preview/examples/LiveTokenGet.json
     */
    /**
     * Sample code: Get live token for resource.
     *
     * @param manager Entry point to ApplicationInsightsManager.
     */
    public static void getLiveTokenForResource(
        com.azure.resourcemanager.applicationinsights.ApplicationInsightsManager manager) {
        manager
            .liveTokens()
            .getWithResponse(
                "subscriptions/df602c9c-7aa0-407d-a6fb-eb20c8bd1192/resourceGroups/FabrikamFiberApp/providers/microsoft.insights/components/CustomAvailabilityTest/providers/microsoft.insights/generatelivetoken",
                Context.NONE);
    }
}
