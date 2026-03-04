
import com.azure.resourcemanager.newrelicobservability.models.TagRule;

/**
 * Samples for TagRules Update.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-05-01-preview/TagRules_Update_MinimumSet_Gen.json
     */
    /**
     * Sample code: TagRules_Update_MinimumSet_Gen.
     * 
     * @param manager Entry point to NewRelicObservabilityManager.
     */
    public static void tagRulesUpdateMinimumSetGen(
        com.azure.resourcemanager.newrelicobservability.NewRelicObservabilityManager manager) {
        TagRule resource = manager.tagRules().getWithResponse("rgopenapi", "ipxmlcbonyxtolzejcjshkmlron",
            "bxcantgzggsepbhqmedjqyrqeezmfb", com.azure.core.util.Context.NONE).getValue();
        resource.update().apply();
    }
}
