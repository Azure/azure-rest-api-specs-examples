
/**
 * Samples for ApplicationTypes List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-01-preview/ApplicationTypeNameListOperation_example.json
     */
    /**
     * Sample code: Get a list of application type name resources.
     * 
     * @param manager Entry point to ServiceFabricManager.
     */
    public static void
        getAListOfApplicationTypeNameResources(com.azure.resourcemanager.servicefabric.ServiceFabricManager manager) {
        manager.applicationTypes().list("resRg", "myCluster", com.azure.core.util.Context.NONE);
    }
}
