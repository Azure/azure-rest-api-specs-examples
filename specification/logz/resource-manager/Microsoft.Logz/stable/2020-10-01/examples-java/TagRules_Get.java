import com.azure.core.util.Context;

/** Samples for TagRules Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/logz/resource-manager/Microsoft.Logz/stable/2020-10-01/examples/TagRules_Get.json
     */
    /**
     * Sample code: TagRules_Get.
     *
     * @param manager Entry point to LogzManager.
     */
    public static void tagRulesGet(com.azure.resourcemanager.logz.LogzManager manager) {
        manager.tagRules().getWithResponse("myResourceGroup", "myMonitor", "default", Context.NONE);
    }
}
