import com.azure.core.util.Context;

/** Samples for SingleSignOn Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/dynatrace/resource-manager/Dynatrace.Observability/stable/2021-09-01/examples/SingleSignOn_Get_MinimumSet_Gen.json
     */
    /**
     * Sample code: SingleSignOn_Get_MinimumSet_Gen.
     *
     * @param manager Entry point to DynatraceManager.
     */
    public static void singleSignOnGetMinimumSetGen(com.azure.resourcemanager.dynatrace.DynatraceManager manager) {
        manager.singleSignOns().getWithResponse("myResourceGroup", "myMonitor", "default", Context.NONE);
    }
}
