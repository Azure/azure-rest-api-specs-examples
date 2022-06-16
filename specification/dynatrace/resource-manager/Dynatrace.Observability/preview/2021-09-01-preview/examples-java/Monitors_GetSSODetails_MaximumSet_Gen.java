import com.azure.core.util.Context;
import com.azure.resourcemanager.dynatrace.models.SsoDetailsRequest;

/** Samples for Monitors GetSsoDetails. */
public final class Main {
    /*
     * x-ms-original-file: specification/dynatrace/resource-manager/Dynatrace.Observability/preview/2021-09-01-preview/examples/Monitors_GetSSODetails_MaximumSet_Gen.json
     */
    /**
     * Sample code: Monitors_GetSSODetails_MaximumSet_Gen.
     *
     * @param manager Entry point to DynatraceManager.
     */
    public static void monitorsGetSSODetailsMaximumSetGen(
        com.azure.resourcemanager.dynatrace.DynatraceManager manager) {
        manager
            .monitors()
            .getSsoDetailsWithResponse(
                "myResourceGroup",
                "myMonitor",
                new SsoDetailsRequest().withUserPrincipal("alice@microsoft.com"),
                Context.NONE);
    }
}
