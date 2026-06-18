
/**
 * Samples for ApplicationTypeVersions List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-01-preview/ApplicationTypeVersionListOperation_example.json
     */
    /**
     * Sample code: Get a list of application type version resources.
     * 
     * @param manager Entry point to ServiceFabricManager.
     */
    public static void getAListOfApplicationTypeVersionResources(
        com.azure.resourcemanager.servicefabric.ServiceFabricManager manager) {
        manager.applicationTypeVersions().list("resRg", "myCluster", "myAppType", com.azure.core.util.Context.NONE);
    }
}
