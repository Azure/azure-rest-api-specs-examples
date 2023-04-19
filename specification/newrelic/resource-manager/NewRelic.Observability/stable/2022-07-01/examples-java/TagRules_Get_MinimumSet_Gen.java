/** Samples for TagRules Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/newrelic/resource-manager/NewRelic.Observability/stable/2022-07-01/examples/TagRules_Get_MinimumSet_Gen.json
     */
    /**
     * Sample code: TagRules_Get_MinimumSet_Gen.
     *
     * @param manager Entry point to NewRelicObservabilityManager.
     */
    public static void tagRulesGetMinimumSetGen(
        com.azure.resourcemanager.newrelicobservability.NewRelicObservabilityManager manager) {
        manager
            .tagRules()
            .getWithResponse(
                "rgopenapi",
                "ipxmlcbonyxtolzejcjshkmlron",
                "bxcantgzggsepbhqmedjqyrqeezmfb",
                com.azure.core.util.Context.NONE);
    }
}
