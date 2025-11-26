
/**
 * Samples for TagRules Delete.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/newrelic/resource-manager/NewRelic.Observability/preview/2025-05-01-preview/examples/
     * TagRules_Delete_MinimumSet_Gen.json
     */
    /**
     * Sample code: TagRules_Delete_MinimumSet_Gen.
     * 
     * @param manager Entry point to NewRelicObservabilityManager.
     */
    public static void tagRulesDeleteMinimumSetGen(
        com.azure.resourcemanager.newrelicobservability.NewRelicObservabilityManager manager) {
        manager.tagRules().delete("rgopenapi", "ipxmlcbonyxtolzejcjshkmlron", "bxcantgzggsepbhqmedjqyrqeezmfb",
            com.azure.core.util.Context.NONE);
    }
}
