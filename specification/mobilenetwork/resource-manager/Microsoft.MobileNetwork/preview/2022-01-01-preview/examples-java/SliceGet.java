import com.azure.core.util.Context;

/** Samples for Slices Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/mobilenetwork/resource-manager/Microsoft.MobileNetwork/preview/2022-01-01-preview/examples/SliceGet.json
     */
    /**
     * Sample code: Get mobile network slice.
     *
     * @param manager Entry point to MobileNetworkManager.
     */
    public static void getMobileNetworkSlice(com.azure.resourcemanager.mobilenetwork.MobileNetworkManager manager) {
        manager.slices().getWithResponse("rg1", "testMobileNetwork", "testSlice", Context.NONE);
    }
}
