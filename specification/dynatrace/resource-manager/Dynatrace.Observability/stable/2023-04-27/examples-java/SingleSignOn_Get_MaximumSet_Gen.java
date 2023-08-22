/** Samples for SingleSignOn Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/dynatrace/resource-manager/Dynatrace.Observability/stable/2023-04-27/examples/SingleSignOn_Get_MaximumSet_Gen.json
     */
    /**
     * Sample code: SingleSignOn_Get_MaximumSet_Gen.
     *
     * @param manager Entry point to DynatraceManager.
     */
    public static void singleSignOnGetMaximumSetGen(com.azure.resourcemanager.dynatrace.DynatraceManager manager) {
        manager
            .singleSignOns()
            .getWithResponse("myResourceGroup", "myMonitor", "default", com.azure.core.util.Context.NONE);
    }
}
