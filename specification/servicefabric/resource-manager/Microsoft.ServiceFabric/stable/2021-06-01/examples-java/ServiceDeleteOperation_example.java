
/**
 * Samples for Services Delete.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/servicefabric/resource-manager/Microsoft.ServiceFabric/stable/2021-06-01/examples/
     * ServiceDeleteOperation_example.json
     */
    /**
     * Sample code: Delete a service.
     * 
     * @param manager Entry point to ServiceFabricManager.
     */
    public static void deleteAService(com.azure.resourcemanager.servicefabric.ServiceFabricManager manager) {
        manager.services().delete("resRg", "myCluster", "myApp", "myService", com.azure.core.util.Context.NONE);
    }
}
