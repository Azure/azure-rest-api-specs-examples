
/**
 * Samples for ApplicationTypes Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-01-preview/ApplicationTypeNameGetOperation_example.json
     */
    /**
     * Sample code: Get an application type.
     * 
     * @param manager Entry point to ServiceFabricManager.
     */
    public static void getAnApplicationType(com.azure.resourcemanager.servicefabric.ServiceFabricManager manager) {
        manager.applicationTypes().getWithResponse("resRg", "myCluster", "myAppType", com.azure.core.util.Context.NONE);
    }
}
