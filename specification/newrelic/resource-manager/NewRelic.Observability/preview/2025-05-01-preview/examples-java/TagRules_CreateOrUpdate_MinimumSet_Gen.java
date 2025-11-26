
/**
 * Samples for TagRules CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/newrelic/resource-manager/NewRelic.Observability/preview/2025-05-01-preview/examples/
     * TagRules_CreateOrUpdate_MinimumSet_Gen.json
     */
    /**
     * Sample code: TagRules_CreateOrUpdate_MinimumSet_Gen.
     * 
     * @param manager Entry point to NewRelicObservabilityManager.
     */
    public static void tagRulesCreateOrUpdateMinimumSetGen(
        com.azure.resourcemanager.newrelicobservability.NewRelicObservabilityManager manager) {
        manager.tagRules().define("bxcantgzggsepbhqmedjqyrqeezmfb")
            .withExistingMonitor("rgopenapi", "ipxmlcbonyxtolzejcjshkmlron").create();
    }
}
