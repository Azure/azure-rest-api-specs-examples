import com.azure.core.util.Context;
import com.azure.resourcemanager.dynatrace.models.TagRule;

/** Samples for TagRules Update. */
public final class Main {
    /*
     * x-ms-original-file: specification/dynatrace/resource-manager/Dynatrace.Observability/preview/2021-09-01-preview/examples/TagRules_Update_MinimumSet_Gen.json
     */
    /**
     * Sample code: TagRules_Update_MinimumSet_Gen.
     *
     * @param manager Entry point to DynatraceManager.
     */
    public static void tagRulesUpdateMinimumSetGen(com.azure.resourcemanager.dynatrace.DynatraceManager manager) {
        TagRule resource =
            manager.tagRules().getWithResponse("myResourceGroup", "myMonitor", "default", Context.NONE).getValue();
        resource.update().apply();
    }
}
