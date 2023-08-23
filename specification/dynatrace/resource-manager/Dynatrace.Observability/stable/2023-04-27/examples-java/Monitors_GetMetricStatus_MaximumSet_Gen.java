/** Samples for Monitors GetMetricStatus. */
public final class Main {
    /*
     * x-ms-original-file: specification/dynatrace/resource-manager/Dynatrace.Observability/stable/2023-04-27/examples/Monitors_GetMetricStatus_MaximumSet_Gen.json
     */
    /**
     * Sample code: Monitors_GetMetricStatus_MaximumSet_Gen.
     *
     * @param manager Entry point to DynatraceManager.
     */
    public static void monitorsGetMetricStatusMaximumSetGen(
        com.azure.resourcemanager.dynatrace.DynatraceManager manager) {
        manager
            .monitors()
            .getMetricStatusWithResponse("rgDynatrace", "fhcjxnxumkdlgpwanewtkdnyuz", com.azure.core.util.Context.NONE);
    }
}
