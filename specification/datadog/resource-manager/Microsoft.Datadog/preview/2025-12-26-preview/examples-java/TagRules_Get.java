
/**
 * Samples for TagRules Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-12-26-preview/TagRules_Get.json
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
