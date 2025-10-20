
/**
 * Samples for TagRules List.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/elastic/resource-manager/Microsoft.Elastic/stable/2025-06-01/examples/TagRules_List.json
     */
    /**
     * Sample code: TagRules_List.
     * 
     * @param manager Entry point to ElasticManager.
     */
    public static void tagRulesList(com.azure.resourcemanager.elastic.ElasticManager manager) {
        manager.tagRules().list("myResourceGroup", "myMonitor", com.azure.core.util.Context.NONE);
    }
}
