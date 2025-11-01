
/**
 * Samples for NamespaceDevices Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-10-01/Get_NamespaceDeviceWithEndpointErrorStatus.json
     */
    /**
     * Sample code: Get NamespaceDevice with Endpoint Error Status.
     * 
     * @param manager Entry point to DeviceRegistryManager.
     */
    public static void getNamespaceDeviceWithEndpointErrorStatus(
        com.azure.resourcemanager.deviceregistry.DeviceRegistryManager manager) {
        manager.namespaceDevices().getWithResponse("myResourceGroup", "my-namespace-1", "my-device-name",
            com.azure.core.util.Context.NONE);
    }
}
