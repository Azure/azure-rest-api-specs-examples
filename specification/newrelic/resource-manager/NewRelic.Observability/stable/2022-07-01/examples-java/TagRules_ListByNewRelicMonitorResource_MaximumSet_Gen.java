/** Samples for TagRules ListByNewRelicMonitorResource. */
public final class Main {
    /*
     * x-ms-original-file: specification/newrelic/resource-manager/NewRelic.Observability/stable/2022-07-01/examples/TagRules_ListByNewRelicMonitorResource_MaximumSet_Gen.json
     */
    /**
     * Sample code: TagRules_ListByNewRelicMonitorResource_MaximumSet_Gen.
     *
     * @param manager Entry point to NewRelicObservabilityManager.
     */
    public static void tagRulesListByNewRelicMonitorResourceMaximumSetGen(
        com.azure.resourcemanager.newrelicobservability.NewRelicObservabilityManager manager) {
        manager
            .tagRules()
            .listByNewRelicMonitorResource(
                "rgopenapi", "ipxmlcbonyxtolzejcjshkmlron", com.azure.core.util.Context.NONE);
    }
}
