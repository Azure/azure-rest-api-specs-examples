
/**
 * Samples for Operations List.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/deviceupdate/resource-manager/Microsoft.DeviceUpdate/stable/2023-07-01/examples/Operations_List.
     * json
     */
    /**
     * Sample code: Gets list of Operations.
     * 
     * @param manager Entry point to DeviceUpdateManager.
     */
    public static void getsListOfOperations(com.azure.resourcemanager.deviceupdate.DeviceUpdateManager manager) {
        manager.operations().list(com.azure.core.util.Context.NONE);
    }
}
