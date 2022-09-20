import com.azure.core.util.Context;
import com.azure.resourcemanager.dynatrace.models.LinkableEnvironmentRequest;

/** Samples for Monitors ListLinkableEnvironments. */
public final class Main {
    /*
     * x-ms-original-file: specification/dynatrace/resource-manager/Dynatrace.Observability/stable/2021-09-01/examples/Monitors_ListLinkableEnvironments_MinimumSet_Gen.json
     */
    /**
     * Sample code: Monitors_ListLinkableEnvironments_MinimumSet_Gen.
     *
     * @param manager Entry point to DynatraceManager.
     */
    public static void monitorsListLinkableEnvironmentsMinimumSetGen(
        com.azure.resourcemanager.dynatrace.DynatraceManager manager) {
        manager
            .monitors()
            .listLinkableEnvironments("myResourceGroup", "myMonitor", new LinkableEnvironmentRequest(), Context.NONE);
    }
}
