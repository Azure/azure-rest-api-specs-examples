import com.azure.core.util.Context;

/** Samples for Slices Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/mobilenetwork/resource-manager/Microsoft.MobileNetwork/preview/2022-04-01-preview/examples/SliceGet.json
     */
    /**
     * Sample code: Get network slice.
     *
     * @param manager Entry point to MobileNetworkManager.
     */
    public static void getNetworkSlice(com.azure.resourcemanager.mobilenetwork.MobileNetworkManager manager) {
        manager.slices().getWithResponse("rg1", "testMobileNetwork", "testSlice", Context.NONE);
    }
}
