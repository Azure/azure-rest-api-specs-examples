import com.azure.core.util.Context;
import com.azure.resourcemanager.dynatrace.models.LinkableEnvironmentRequest;

/** Samples for Monitors ListLinkableEnvironments. */
public final class Main {
    /*
     * x-ms-original-file: specification/dynatrace/resource-manager/Dynatrace.Observability/preview/2021-09-01-preview/examples/Monitors_ListLinkableEnvironments_MaximumSet_Gen.json
     */
    /**
     * Sample code: Monitors_ListLinkableEnvironments_MaximumSet_Gen.
     *
     * @param manager Entry point to DynatraceManager.
     */
    public static void monitorsListLinkableEnvironmentsMaximumSetGen(
        com.azure.resourcemanager.dynatrace.DynatraceManager manager) {
        manager
            .monitors()
            .listLinkableEnvironments(
                "myResourceGroup",
                "myMonitor",
                new LinkableEnvironmentRequest()
                    .withTenantId("00000000-0000-0000-0000-000000000000")
                    .withUserPrincipal("alice@microsoft.com")
                    .withRegion("East US"),
                Context.NONE);
    }
}
