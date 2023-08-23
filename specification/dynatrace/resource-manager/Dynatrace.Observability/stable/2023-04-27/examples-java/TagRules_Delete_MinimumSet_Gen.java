/** Samples for TagRules Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/dynatrace/resource-manager/Dynatrace.Observability/stable/2023-04-27/examples/TagRules_Delete_MinimumSet_Gen.json
     */
    /**
     * Sample code: TagRules_Delete_MinimumSet_Gen.
     *
     * @param manager Entry point to DynatraceManager.
     */
    public static void tagRulesDeleteMinimumSetGen(com.azure.resourcemanager.dynatrace.DynatraceManager manager) {
        manager.tagRules().delete("myResourceGroup", "myMonitor", "default", com.azure.core.util.Context.NONE);
    }
}
