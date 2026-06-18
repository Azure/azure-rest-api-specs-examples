
/**
 * Samples for ApplicationTypeVersions Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-01-preview/ApplicationTypeVersionGetOperation_example.json
     */
    /**
     * Sample code: Get an application type version.
     * 
     * @param manager Entry point to ServiceFabricManager.
     */
    public static void
        getAnApplicationTypeVersion(com.azure.resourcemanager.servicefabric.ServiceFabricManager manager) {
        manager.applicationTypeVersions().getWithResponse("resRg", "myCluster", "myAppType", "1.0",
            com.azure.core.util.Context.NONE);
    }
}
