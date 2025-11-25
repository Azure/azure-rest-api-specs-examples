
import com.azure.resourcemanager.newrelicobservability.models.AppServicesGetRequest;
import java.util.Arrays;

/**
 * Samples for Monitors ListAppServices.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/newrelic/resource-manager/NewRelic.Observability/preview/2025-05-01-preview/examples/
     * Monitors_ListAppServices_MinimumSet_Gen.json
     */
    /**
     * Sample code: Monitors_ListAppServices_MinimumSet_Gen.
     * 
     * @param manager Entry point to NewRelicObservabilityManager.
     */
    public static void monitorsListAppServicesMinimumSetGen(
        com.azure.resourcemanager.newrelicobservability.NewRelicObservabilityManager manager) {
        manager.monitors().listAppServices("rgNewRelic", "fhcjxnxumkdlgpwanewtkdnyuz",
            new AppServicesGetRequest().withAzureResourceIds(Arrays.asList(
                "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rgNewRelic/providers/NewRelic.Observability/monitors/fhcjxnxumkdlgpwanewtkdnyuz"))
                .withUserEmail("ruxvg@xqkmdhrnoo.hlmbpm"),
            com.azure.core.util.Context.NONE);
    }
}
