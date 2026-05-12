
/**
 * Samples for RegistryEndpoint ListByInstanceResource.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-01/RegistryEndpoint_ListByInstanceResource_MaximumSet_Gen.json
     */
    /**
     * Sample code: RegistryEndpoint_ListByInstanceResource_MaximumSet.
     * 
     * @param manager Entry point to IoTOperationsManager.
     */
    public static void registryEndpointListByInstanceResourceMaximumSet(
        com.azure.resourcemanager.iotoperations.IoTOperationsManager manager) {
        manager.registryEndpoints().listByInstanceResource("rgiotoperations", "resource-123",
            com.azure.core.util.Context.NONE);
    }
}
