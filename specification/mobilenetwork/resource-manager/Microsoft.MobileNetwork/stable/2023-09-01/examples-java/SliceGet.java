/** Samples for Slices Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/mobilenetwork/resource-manager/Microsoft.MobileNetwork/stable/2023-09-01/examples/SliceGet.json
     */
    /**
     * Sample code: Get network slice.
     *
     * @param manager Entry point to MobileNetworkManager.
     */
    public static void getNetworkSlice(com.azure.resourcemanager.mobilenetwork.MobileNetworkManager manager) {
        manager.slices().getWithResponse("rg1", "testMobileNetwork", "testSlice", com.azure.core.util.Context.NONE);
    }
}
