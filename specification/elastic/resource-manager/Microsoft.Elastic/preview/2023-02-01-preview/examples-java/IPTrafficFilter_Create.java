/** Samples for CreateAndAssociateIpFilter Create. */
public final class Main {
    /*
     * x-ms-original-file: specification/elastic/resource-manager/Microsoft.Elastic/preview/2023-02-01-preview/examples/IPTrafficFilter_Create.json
     */
    /**
     * Sample code: createAndAssociateIPFilter_Create.
     *
     * @param manager Entry point to ElasticManager.
     */
    public static void createAndAssociateIPFilterCreate(com.azure.resourcemanager.elastic.ElasticManager manager) {
        manager
            .createAndAssociateIpFilters()
            .create(
                "myResourceGroup",
                "myMonitor",
                "192.168.131.0, 192.168.132.6/22",
                null,
                com.azure.core.util.Context.NONE);
    }
}
