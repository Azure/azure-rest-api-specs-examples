import com.azure.resourcemanager.mobilenetwork.models.Snssai;

/** Samples for Slices CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/mobilenetwork/resource-manager/Microsoft.MobileNetwork/preview/2022-01-01-preview/examples/SliceCreate.json
     */
    /**
     * Sample code: Create mobile network slice.
     *
     * @param manager Entry point to MobileNetworkManager.
     */
    public static void createMobileNetworkSlice(com.azure.resourcemanager.mobilenetwork.MobileNetworkManager manager) {
        manager
            .slices()
            .define("testSlice")
            .withRegion("eastus")
            .withExistingMobileNetwork("rg1", "testMobileNetwork")
            .withSnssai(new Snssai().withSst(1).withSd("1abcde"))
            .withDescription("myFavouriteSlice")
            .create();
    }
}
