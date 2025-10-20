
/**
 * Samples for TagRules CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/elastic/resource-manager/Microsoft.Elastic/stable/2025-06-01/examples/TagRules_CreateOrUpdate.json
     */
    /**
     * Sample code: TagRules_CreateOrUpdate.
     * 
     * @param manager Entry point to ElasticManager.
     */
    public static void tagRulesCreateOrUpdate(com.azure.resourcemanager.elastic.ElasticManager manager) {
        manager.tagRules().define("default").withExistingMonitor("myResourceGroup", "myMonitor").create();
    }
}
