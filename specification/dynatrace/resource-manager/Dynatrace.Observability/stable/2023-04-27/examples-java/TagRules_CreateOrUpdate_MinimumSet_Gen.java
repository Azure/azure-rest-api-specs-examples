/** Samples for TagRules CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/dynatrace/resource-manager/Dynatrace.Observability/stable/2023-04-27/examples/TagRules_CreateOrUpdate_MinimumSet_Gen.json
     */
    /**
     * Sample code: TagRules_CreateOrUpdate_MinimumSet_Gen.
     *
     * @param manager Entry point to DynatraceManager.
     */
    public static void tagRulesCreateOrUpdateMinimumSetGen(
        com.azure.resourcemanager.dynatrace.DynatraceManager manager) {
        manager.tagRules().define("default").withExistingMonitor("myResourceGroup", "myMonitor").create();
    }
}
