import com.azure.core.util.Context;

/** Samples for TagRules Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/logz/resource-manager/Microsoft.Logz/stable/2020-10-01/examples/TagRules_Delete.json
     */
    /**
     * Sample code: TagRules_Delete.
     *
     * @param manager Entry point to LogzManager.
     */
    public static void tagRulesDelete(com.azure.resourcemanager.logz.LogzManager manager) {
        manager.tagRules().deleteWithResponse("myResourceGroup", "myMonitor", "default", Context.NONE);
    }
}
