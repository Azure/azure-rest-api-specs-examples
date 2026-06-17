
/**
 * Samples for Services List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-01-preview/ServiceListOperation_example.json
     */
    /**
     * Sample code: Get a list of service resources.
     * 
     * @param manager Entry point to ServiceFabricManager.
     */
    public static void
        getAListOfServiceResources(com.azure.resourcemanager.servicefabric.ServiceFabricManager manager) {
        manager.services().list("resRg", "myCluster", "myApp", com.azure.core.util.Context.NONE);
    }
}
