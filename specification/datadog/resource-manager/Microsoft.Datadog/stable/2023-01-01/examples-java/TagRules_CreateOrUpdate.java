/** Samples for TagRules CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/datadog/resource-manager/Microsoft.Datadog/stable/2023-01-01/examples/TagRules_CreateOrUpdate.json
     */
    /**
     * Sample code: TagRules_CreateOrUpdate.
     *
     * @param manager Entry point to MicrosoftDatadogManager.
     */
    public static void tagRulesCreateOrUpdate(com.azure.resourcemanager.datadog.MicrosoftDatadogManager manager) {
        manager.tagRules().define("default").withExistingMonitor("myResourceGroup", "myMonitor").create();
    }
}
