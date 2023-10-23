/** Samples for TagRules Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/datadog/resource-manager/Microsoft.Datadog/stable/2023-01-01/examples/TagRules_Get.json
     */
    /**
     * Sample code: TagRules_Get.
     *
     * @param manager Entry point to MicrosoftDatadogManager.
     */
    public static void tagRulesGet(com.azure.resourcemanager.datadog.MicrosoftDatadogManager manager) {
        manager.tagRules().getWithResponse("myResourceGroup", "myMonitor", "default", com.azure.core.util.Context.NONE);
    }
}
