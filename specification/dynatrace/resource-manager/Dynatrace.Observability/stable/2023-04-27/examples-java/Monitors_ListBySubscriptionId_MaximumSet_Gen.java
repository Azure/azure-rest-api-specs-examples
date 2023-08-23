/** Samples for Monitors List. */
public final class Main {
    /*
     * x-ms-original-file: specification/dynatrace/resource-manager/Dynatrace.Observability/stable/2023-04-27/examples/Monitors_ListBySubscriptionId_MaximumSet_Gen.json
     */
    /**
     * Sample code: Monitors_ListBySubscriptionId_MaximumSet_Gen.
     *
     * @param manager Entry point to DynatraceManager.
     */
    public static void monitorsListBySubscriptionIdMaximumSetGen(
        com.azure.resourcemanager.dynatrace.DynatraceManager manager) {
        manager.monitors().list(com.azure.core.util.Context.NONE);
    }
}
