
/**
 * Samples for TagRules Delete.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/elastic/resource-manager/Microsoft.Elastic/stable/2024-03-01/examples/TagRules_Delete.json
     */
    /**
     * Sample code: TagRules_Delete.
     * 
     * @param manager Entry point to ElasticManager.
     */
    public static void tagRulesDelete(com.azure.resourcemanager.elastic.ElasticManager manager) {
        manager.tagRules().delete("myResourceGroup", "myMonitor", "default", com.azure.core.util.Context.NONE);
    }
}
