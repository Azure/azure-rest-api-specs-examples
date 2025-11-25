
import com.azure.resourcemanager.newrelicobservability.models.SwitchBillingRequest;

/**
 * Samples for Monitors SwitchBilling.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/newrelic/resource-manager/NewRelic.Observability/preview/2025-05-01-preview/examples/
     * Monitors_SwitchBilling_MinimumSet_Gen.json
     */
    /**
     * Sample code: Monitors_SwitchBilling_MinimumSet_Gen.
     * 
     * @param manager Entry point to NewRelicObservabilityManager.
     */
    public static void monitorsSwitchBillingMinimumSetGen(
        com.azure.resourcemanager.newrelicobservability.NewRelicObservabilityManager manager) {
        manager.monitors().switchBillingWithResponse("rgNewRelic", "fhcjxnxumkdlgpwanewtkdnyuz",
            new SwitchBillingRequest().withUserEmail("ruxvg@xqkmdhrnoo.hlmbpm"), com.azure.core.util.Context.NONE);
    }
}
