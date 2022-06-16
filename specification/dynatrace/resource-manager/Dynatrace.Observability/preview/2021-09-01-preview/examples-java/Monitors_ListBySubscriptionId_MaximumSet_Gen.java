import com.azure.core.util.Context;

/** Samples for Monitors List. */
public final class Main {
    /*
     * x-ms-original-file: specification/dynatrace/resource-manager/Dynatrace.Observability/preview/2021-09-01-preview/examples/Monitors_ListBySubscriptionId_MaximumSet_Gen.json
     */
    /**
     * Sample code: Monitors_ListBySubscriptionId_MaximumSet_Gen.
     *
     * @param manager Entry point to DynatraceManager.
     */
    public static void monitorsListBySubscriptionIdMaximumSetGen(
        com.azure.resourcemanager.dynatrace.DynatraceManager manager) {
        manager.monitors().list(Context.NONE);
    }
}
