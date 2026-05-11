
/**
 * Samples for RegistryEndpoint Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-01/RegistryEndpoint_Delete_MaximumSet_Gen.json
     */
    /**
     * Sample code: RegistryEndpoint_Delete_MaximumSet.
     * 
     * @param manager Entry point to IoTOperationsManager.
     */
    public static void
        registryEndpointDeleteMaximumSet(com.azure.resourcemanager.iotoperations.IoTOperationsManager manager) {
        manager.registryEndpoints().delete("rgiotoperations", "resource-123", "resource-123",
            com.azure.core.util.Context.NONE);
    }
}
