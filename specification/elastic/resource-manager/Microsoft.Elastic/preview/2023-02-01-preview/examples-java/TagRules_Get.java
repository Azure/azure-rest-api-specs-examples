/** Samples for TagRules Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/elastic/resource-manager/Microsoft.Elastic/preview/2023-02-01-preview/examples/TagRules_Get.json
     */
    /**
     * Sample code: TagRules_Get.
     *
     * @param manager Entry point to ElasticManager.
     */
    public static void tagRulesGet(com.azure.resourcemanager.elastic.ElasticManager manager) {
        manager.tagRules().getWithResponse("myResourceGroup", "myMonitor", "default", com.azure.core.util.Context.NONE);
    }
}
