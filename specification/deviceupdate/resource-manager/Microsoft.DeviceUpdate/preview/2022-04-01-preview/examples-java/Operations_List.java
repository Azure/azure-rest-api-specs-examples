import com.azure.core.util.Context;

/** Samples for Operations List. */
public final class Main {
    /*
     * x-ms-original-file: specification/deviceupdate/resource-manager/Microsoft.DeviceUpdate/preview/2022-04-01-preview/examples/Operations_List.json
     */
    /**
     * Sample code: Gets list of Operations.
     *
     * @param manager Entry point to DeviceUpdateManager.
     */
    public static void getsListOfOperations(com.azure.resourcemanager.deviceupdate.DeviceUpdateManager manager) {
        manager.operations().list(Context.NONE);
    }
}
