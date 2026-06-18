
/**
 * Samples for ApplicationTypes Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-01-preview/ApplicationTypeNameDeleteOperation_example.json
     */
    /**
     * Sample code: Delete an application type.
     * 
     * @param manager Entry point to ServiceFabricManager.
     */
    public static void deleteAnApplicationType(com.azure.resourcemanager.servicefabric.ServiceFabricManager manager) {
        manager.applicationTypes().delete("resRg", "myCluster", "myAppType", com.azure.core.util.Context.NONE);
    }
}
