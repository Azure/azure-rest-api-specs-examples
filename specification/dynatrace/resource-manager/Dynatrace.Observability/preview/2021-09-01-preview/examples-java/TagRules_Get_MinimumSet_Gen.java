import com.azure.core.util.Context;

/** Samples for TagRules Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/dynatrace/resource-manager/Dynatrace.Observability/preview/2021-09-01-preview/examples/TagRules_Get_MinimumSet_Gen.json
     */
    /**
     * Sample code: TagRules_Get_MinimumSet_Gen.
     *
     * @param manager Entry point to DynatraceManager.
     */
    public static void tagRulesGetMinimumSetGen(com.azure.resourcemanager.dynatrace.DynatraceManager manager) {
        manager.tagRules().getWithResponse("myResourceGroup", "myMonitor", "default", Context.NONE);
    }
}
