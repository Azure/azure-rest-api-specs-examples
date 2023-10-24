/** Samples for TagRules List. */
public final class Main {
    /*
     * x-ms-original-file: specification/datadog/resource-manager/Microsoft.Datadog/stable/2023-01-01/examples/TagRules_List.json
     */
    /**
     * Sample code: TagRules_List.
     *
     * @param manager Entry point to MicrosoftDatadogManager.
     */
    public static void tagRulesList(com.azure.resourcemanager.datadog.MicrosoftDatadogManager manager) {
        manager.tagRules().list("myResourceGroup", "myMonitor", com.azure.core.util.Context.NONE);
    }
}
