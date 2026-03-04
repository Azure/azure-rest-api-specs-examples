
/**
 * Samples for ConnectedPartnerResources List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-06-01/ConnectedPartnerResources_List.json
     */
    /**
     * Sample code: ConnectedPartnerResources_List.
     * 
     * @param manager Entry point to ElasticManager.
     */
    public static void connectedPartnerResourcesList(com.azure.resourcemanager.elastic.ElasticManager manager) {
        manager.connectedPartnerResources().list("myResourceGroup", "myMonitor", com.azure.core.util.Context.NONE);
    }
}
