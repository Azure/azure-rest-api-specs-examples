import com.azure.resourcemanager.newrelicobservability.models.HostsGetRequest;
import java.util.Arrays;

/** Samples for Monitors ListHosts. */
public final class Main {
    /*
     * x-ms-original-file: specification/newrelic/resource-manager/NewRelic.Observability/stable/2022-07-01/examples/Monitors_ListHosts_MaximumSet_Gen.json
     */
    /**
     * Sample code: Monitors_ListHosts_MaximumSet_Gen.
     *
     * @param manager Entry point to NewRelicObservabilityManager.
     */
    public static void monitorsListHostsMaximumSetGen(
        com.azure.resourcemanager.newrelicobservability.NewRelicObservabilityManager manager) {
        manager
            .monitors()
            .listHosts(
                "rgopenapi",
                "ipxmlcbonyxtolzejcjshkmlron",
                new HostsGetRequest()
                    .withVmIds(Arrays.asList("xzphvxvfmvjrnsgyns"))
                    .withUserEmail("ruxvg@xqkmdhrnoo.hlmbpm"),
                com.azure.core.util.Context.NONE);
    }
}
