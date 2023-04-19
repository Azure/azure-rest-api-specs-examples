import com.azure.resourcemanager.newrelicobservability.models.AppServicesGetRequest;

/** Samples for Monitors ListAppServices. */
public final class Main {
    /*
     * x-ms-original-file: specification/newrelic/resource-manager/NewRelic.Observability/stable/2022-07-01/examples/Monitors_ListAppServices_MinimumSet_Gen.json
     */
    /**
     * Sample code: Monitors_ListAppServices_MinimumSet_Gen.
     *
     * @param manager Entry point to NewRelicObservabilityManager.
     */
    public static void monitorsListAppServicesMinimumSetGen(
        com.azure.resourcemanager.newrelicobservability.NewRelicObservabilityManager manager) {
        manager
            .monitors()
            .listAppServices(
                "rgNewRelic",
                "fhcjxnxumkdlgpwanewtkdnyuz",
                new AppServicesGetRequest().withUserEmail("ruxvg@xqkmdhrnoo.hlmbpm"),
                com.azure.core.util.Context.NONE);
    }
}
