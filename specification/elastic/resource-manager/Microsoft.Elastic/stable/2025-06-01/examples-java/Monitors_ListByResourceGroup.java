
/**
 * Samples for Monitors ListByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/elastic/resource-manager/Microsoft.Elastic/stable/2025-06-01/examples/Monitors_ListByResourceGroup.
     * json
     */
    /**
     * Sample code: Monitors_ListByResourceGroup.
     * 
     * @param manager Entry point to ElasticManager.
     */
    public static void monitorsListByResourceGroup(com.azure.resourcemanager.elastic.ElasticManager manager) {
        manager.monitors().listByResourceGroup("myResourceGroup", com.azure.core.util.Context.NONE);
    }
}
