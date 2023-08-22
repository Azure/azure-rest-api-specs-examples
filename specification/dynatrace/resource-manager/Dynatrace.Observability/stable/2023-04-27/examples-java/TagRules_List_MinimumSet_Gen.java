/** Samples for TagRules List. */
public final class Main {
    /*
     * x-ms-original-file: specification/dynatrace/resource-manager/Dynatrace.Observability/stable/2023-04-27/examples/TagRules_List_MinimumSet_Gen.json
     */
    /**
     * Sample code: TagRules_List_MinimumSet_Gen.
     *
     * @param manager Entry point to DynatraceManager.
     */
    public static void tagRulesListMinimumSetGen(com.azure.resourcemanager.dynatrace.DynatraceManager manager) {
        manager.tagRules().list("myResourceGroup", "myMonitor", com.azure.core.util.Context.NONE);
    }
}
