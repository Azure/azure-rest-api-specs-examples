import com.azure.core.util.Context;

/** Samples for SingleSignOn Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/dynatrace/resource-manager/Dynatrace.Observability/preview/2021-09-01-preview/examples/SingleSignOn_Get_MaximumSet_Gen.json
     */
    /**
     * Sample code: SingleSignOn_Get_MaximumSet_Gen.
     *
     * @param manager Entry point to DynatraceManager.
     */
    public static void singleSignOnGetMaximumSetGen(com.azure.resourcemanager.dynatrace.DynatraceManager manager) {
        manager.singleSignOns().getWithResponse("myResourceGroup", "myMonitor", "default", Context.NONE);
    }
}
