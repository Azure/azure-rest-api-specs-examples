import com.azure.core.util.Context;

/** Samples for TagRules List. */
public final class Main {
    /*
     * x-ms-original-file: specification/dynatrace/resource-manager/Dynatrace.Observability/preview/2021-09-01-preview/examples/TagRules_List_MinimumSet_Gen.json
     */
    /**
     * Sample code: TagRules_List_MinimumSet_Gen.
     *
     * @param manager Entry point to DynatraceManager.
     */
    public static void tagRulesListMinimumSetGen(com.azure.resourcemanager.dynatrace.DynatraceManager manager) {
        manager.tagRules().list("myResourceGroup", "myMonitor", Context.NONE);
    }
}
