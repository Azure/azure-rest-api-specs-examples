/** Samples for Monitors List. */
public final class Main {
    /*
     * x-ms-original-file: specification/dynatrace/resource-manager/Dynatrace.Observability/stable/2023-04-27/examples/Monitors_ListBySubscriptionId_MinimumSet_Gen.json
     */
    /**
     * Sample code: Monitors_ListBySubscriptionId_MinimumSet_Gen.
     *
     * @param manager Entry point to DynatraceManager.
     */
    public static void monitorsListBySubscriptionIdMinimumSetGen(
        com.azure.resourcemanager.dynatrace.DynatraceManager manager) {
        manager.monitors().list(com.azure.core.util.Context.NONE);
    }
}
