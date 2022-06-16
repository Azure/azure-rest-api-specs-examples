import com.azure.core.util.Context;

/** Samples for TagRules Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/dynatrace/resource-manager/Dynatrace.Observability/preview/2021-09-01-preview/examples/TagRules_Delete_MaximumSet_Gen.json
     */
    /**
     * Sample code: TagRules_Delete_MaximumSet_Gen.
     *
     * @param manager Entry point to DynatraceManager.
     */
    public static void tagRulesDeleteMaximumSetGen(com.azure.resourcemanager.dynatrace.DynatraceManager manager) {
        manager.tagRules().delete("myResourceGroup", "myMonitor", "default", Context.NONE);
    }
}
