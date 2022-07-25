import com.azure.core.util.Context;

/** Samples for Slices Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/mobilenetwork/resource-manager/Microsoft.MobileNetwork/preview/2022-04-01-preview/examples/SliceDelete.json
     */
    /**
     * Sample code: Delete network slice.
     *
     * @param manager Entry point to MobileNetworkManager.
     */
    public static void deleteNetworkSlice(com.azure.resourcemanager.mobilenetwork.MobileNetworkManager manager) {
        manager.slices().delete("rg1", "testMobileNetwork", "testSlice", Context.NONE);
    }
}
