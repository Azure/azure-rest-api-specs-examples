/** Samples for Operations List. */
public final class Main {
    /*
     * x-ms-original-file: specification/iothub/resource-manager/Microsoft.Devices/preview/2022-11-15-preview/examples/iothub_operations.json
     */
    /**
     * Sample code: Operations_List.
     *
     * @param manager Entry point to IotHubManager.
     */
    public static void operationsList(com.azure.resourcemanager.iothub.IotHubManager manager) {
        manager.operations().list(com.azure.core.util.Context.NONE);
    }
}
