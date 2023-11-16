/** Samples for Slices Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/mobilenetwork/resource-manager/Microsoft.MobileNetwork/stable/2023-09-01/examples/SliceDelete.json
     */
    /**
     * Sample code: Delete network slice.
     *
     * @param manager Entry point to MobileNetworkManager.
     */
    public static void deleteNetworkSlice(com.azure.resourcemanager.mobilenetwork.MobileNetworkManager manager) {
        manager.slices().delete("rg1", "testMobileNetwork", "testSlice", com.azure.core.util.Context.NONE);
    }
}
