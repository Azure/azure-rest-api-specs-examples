/** Samples for TagRules List. */
public final class Main {
    /*
     * x-ms-original-file: specification/logz/resource-manager/Microsoft.Logz/stable/2020-10-01/examples/TagRules_List.json
     */
    /**
     * Sample code: TagRules_List.
     *
     * @param manager Entry point to LogzManager.
     */
    public static void tagRulesList(com.azure.resourcemanager.logz.LogzManager manager) {
        manager.tagRules().list("myResourceGroup", "myMonitor", com.azure.core.util.Context.NONE);
    }
}
