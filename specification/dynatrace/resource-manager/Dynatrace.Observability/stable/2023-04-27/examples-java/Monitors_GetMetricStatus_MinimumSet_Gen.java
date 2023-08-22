/** Samples for Monitors GetMetricStatus. */
public final class Main {
    /*
     * x-ms-original-file: specification/dynatrace/resource-manager/Dynatrace.Observability/stable/2023-04-27/examples/Monitors_GetMetricStatus_MinimumSet_Gen.json
     */
    /**
     * Sample code: Monitors_GetMetricStatus_MinimumSet_Gen.
     *
     * @param manager Entry point to DynatraceManager.
     */
    public static void monitorsGetMetricStatusMinimumSetGen(
        com.azure.resourcemanager.dynatrace.DynatraceManager manager) {
        manager
            .monitors()
            .getMetricStatusWithResponse("rgDynatrace", "fhcjxnxumkdlgpwanewtkdnyuz", com.azure.core.util.Context.NONE);
    }
}
