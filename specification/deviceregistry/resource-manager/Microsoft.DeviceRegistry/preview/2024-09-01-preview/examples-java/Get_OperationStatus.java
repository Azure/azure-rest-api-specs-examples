
/**
 * Samples for OperationStatus Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2024-09-01-preview/Get_OperationStatus.json
     */
    /**
     * Sample code: Get_OperationStatus.
     * 
     * @param manager Entry point to DeviceRegistryManager.
     */
    public static void getOperationStatus(com.azure.resourcemanager.deviceregistry.DeviceRegistryManager manager) {
        manager.operationStatus().getWithResponse("testLocation", "xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx",
            com.azure.core.util.Context.NONE);
    }
}
