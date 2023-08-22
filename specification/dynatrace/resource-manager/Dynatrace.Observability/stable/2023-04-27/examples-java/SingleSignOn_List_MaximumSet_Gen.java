/** Samples for SingleSignOn List. */
public final class Main {
    /*
     * x-ms-original-file: specification/dynatrace/resource-manager/Dynatrace.Observability/stable/2023-04-27/examples/SingleSignOn_List_MaximumSet_Gen.json
     */
    /**
     * Sample code: SingleSignOn_List_MaximumSet_Gen.
     *
     * @param manager Entry point to DynatraceManager.
     */
    public static void singleSignOnListMaximumSetGen(com.azure.resourcemanager.dynatrace.DynatraceManager manager) {
        manager.singleSignOns().list("myResourceGroup", "myMonitor", com.azure.core.util.Context.NONE);
    }
}
