
/**
 * Samples for RegistryEndpoint Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-01/RegistryEndpoint_Get_MaximumSet_Gen.json
     */
    /**
     * Sample code: RegistryEndpoint_Get_MaximumSet.
     * 
     * @param manager Entry point to IoTOperationsManager.
     */
    public static void
        registryEndpointGetMaximumSet(com.azure.resourcemanager.iotoperations.IoTOperationsManager manager) {
        manager.registryEndpoints().getWithResponse("rgiotoperations", "resource-123", "resource-123",
            com.azure.core.util.Context.NONE);
    }
}
