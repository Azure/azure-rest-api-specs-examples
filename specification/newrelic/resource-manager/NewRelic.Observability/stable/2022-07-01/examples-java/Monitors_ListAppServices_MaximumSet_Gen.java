import com.azure.resourcemanager.newrelicobservability.models.AppServicesGetRequest;
import java.util.Arrays;

/** Samples for Monitors ListAppServices. */
public final class Main {
    /*
     * x-ms-original-file: specification/newrelic/resource-manager/NewRelic.Observability/stable/2022-07-01/examples/Monitors_ListAppServices_MaximumSet_Gen.json
     */
    /**
     * Sample code: Monitors_ListAppServices_MaximumSet_Gen.
     *
     * @param manager Entry point to NewRelicObservabilityManager.
     */
    public static void monitorsListAppServicesMaximumSetGen(
        com.azure.resourcemanager.newrelicobservability.NewRelicObservabilityManager manager) {
        manager
            .monitors()
            .listAppServices(
                "rgNewRelic",
                "fhcjxnxumkdlgpwanewtkdnyuz",
                new AppServicesGetRequest()
                    .withAzureResourceIds(Arrays.asList("pvzrksrmzowobuhxpwiotnpcvjbu"))
                    .withUserEmail("ruxvg@xqkmdhrnoo.hlmbpm"),
                com.azure.core.util.Context.NONE);
    }
}
