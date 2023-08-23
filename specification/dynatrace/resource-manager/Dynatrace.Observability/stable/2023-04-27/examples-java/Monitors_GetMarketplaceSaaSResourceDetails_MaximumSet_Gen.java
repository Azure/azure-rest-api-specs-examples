import com.azure.resourcemanager.dynatrace.models.MarketplaceSaaSResourceDetailsRequest;

/** Samples for Monitors GetMarketplaceSaaSResourceDetails. */
public final class Main {
    /*
     * x-ms-original-file: specification/dynatrace/resource-manager/Dynatrace.Observability/stable/2023-04-27/examples/Monitors_GetMarketplaceSaaSResourceDetails_MaximumSet_Gen.json
     */
    /**
     * Sample code: Monitors_GetMarketplaceSaaSResourceDetails_MaximumSet_Gen.
     *
     * @param manager Entry point to DynatraceManager.
     */
    public static void monitorsGetMarketplaceSaaSResourceDetailsMaximumSetGen(
        com.azure.resourcemanager.dynatrace.DynatraceManager manager) {
        manager
            .monitors()
            .getMarketplaceSaaSResourceDetailsWithResponse(
                new MarketplaceSaaSResourceDetailsRequest().withTenantId("urnmattojzhktcfw"),
                com.azure.core.util.Context.NONE);
    }
}
