import com.azure.core.util.Context;

/** Samples for Operations List. */
public final class Main {
    /*
     * x-ms-original-file: specification/iothub/resource-manager/Microsoft.Devices/stable/2021-07-02/examples/iothub_operations.json
     */
    /**
     * Sample code: Operations_List.
     *
     * @param manager Entry point to IotHubManager.
     */
    public static void operationsList(com.azure.resourcemanager.iothub.IotHubManager manager) {
        manager.operations().list(Context.NONE);
    }
}
