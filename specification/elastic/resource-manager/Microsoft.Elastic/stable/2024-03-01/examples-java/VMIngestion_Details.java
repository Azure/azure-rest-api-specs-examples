
/**
 * Samples for VMIngestion Details.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/elastic/resource-manager/Microsoft.Elastic/stable/2024-03-01/examples/VMIngestion_Details.json
     */
    /**
     * Sample code: VMIngestion_Details.
     * 
     * @param manager Entry point to ElasticManager.
     */
    public static void vMIngestionDetails(com.azure.resourcemanager.elastic.ElasticManager manager) {
        manager.vMIngestions().detailsWithResponse("myResourceGroup", "myMonitor", com.azure.core.util.Context.NONE);
    }
}
