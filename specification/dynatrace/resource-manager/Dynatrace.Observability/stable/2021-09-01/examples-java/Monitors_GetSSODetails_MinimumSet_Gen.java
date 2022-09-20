import com.azure.core.util.Context;
import com.azure.resourcemanager.dynatrace.models.SsoDetailsRequest;

/** Samples for Monitors GetSsoDetails. */
public final class Main {
    /*
     * x-ms-original-file: specification/dynatrace/resource-manager/Dynatrace.Observability/stable/2021-09-01/examples/Monitors_GetSSODetails_MinimumSet_Gen.json
     */
    /**
     * Sample code: Monitors_GetSSODetails_MinimumSet_Gen.
     *
     * @param manager Entry point to DynatraceManager.
     */
    public static void monitorsGetSSODetailsMinimumSetGen(
        com.azure.resourcemanager.dynatrace.DynatraceManager manager) {
        manager
            .monitors()
            .getSsoDetailsWithResponse("myResourceGroup", "myMonitor", new SsoDetailsRequest(), Context.NONE);
    }
}
