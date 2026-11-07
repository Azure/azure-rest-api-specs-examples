
/**
 * Samples for DataCollectionRuleAssociations ListByDataCollectionEndpoint.
 */
public final class Main {
    /*
     * x-ms-original-file: 2024-03-11/DataCollectionRuleAssociationsListByDataCollectionEndpoint.json
     */
    /**
     * Sample code: List associations for specified data collection endpoint.
     * 
     * @param manager Entry point to MonitorManager.
     */
    public static void
        listAssociationsForSpecifiedDataCollectionEndpoint(com.azure.resourcemanager.monitor.MonitorManager manager) {
        manager.serviceClient().getDataCollectionRuleAssociations().listByDataCollectionEndpoint("myResourceGroup",
            "myDataCollectionEndpointName", com.azure.core.util.Context.NONE);
    }
}
