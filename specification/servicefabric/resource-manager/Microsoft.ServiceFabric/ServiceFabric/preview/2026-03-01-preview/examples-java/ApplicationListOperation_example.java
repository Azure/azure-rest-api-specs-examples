
/**
 * Samples for Applications List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-01-preview/ApplicationListOperation_example.json
     */
    /**
     * Sample code: Get a list of application resources.
     * 
     * @param manager Entry point to ServiceFabricManager.
     */
    public static void
        getAListOfApplicationResources(com.azure.resourcemanager.servicefabric.ServiceFabricManager manager) {
        manager.applications().list("resRg", "myCluster", com.azure.core.util.Context.NONE);
    }
}
