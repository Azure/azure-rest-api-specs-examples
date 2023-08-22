/** Samples for TagRules Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/dynatrace/resource-manager/Dynatrace.Observability/stable/2023-04-27/examples/TagRules_Get_MaximumSet_Gen.json
     */
    /**
     * Sample code: TagRules_Get_MaximumSet_Gen.
     *
     * @param manager Entry point to DynatraceManager.
     */
    public static void tagRulesGetMaximumSetGen(com.azure.resourcemanager.dynatrace.DynatraceManager manager) {
        manager.tagRules().getWithResponse("myResourceGroup", "myMonitor", "default", com.azure.core.util.Context.NONE);
    }
}
