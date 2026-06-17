
/**
 * Samples for Applications Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-01-preview/ApplicationDeleteOperation_example.json
     */
    /**
     * Sample code: Delete an application.
     * 
     * @param manager Entry point to ServiceFabricManager.
     */
    public static void deleteAnApplication(com.azure.resourcemanager.servicefabric.ServiceFabricManager manager) {
        manager.applications().delete("resRg", "myCluster", "myApp", com.azure.core.util.Context.NONE);
    }
}
