import com.azure.resourcemanager.mobilenetwork.models.Snssai;

/** Samples for Slices CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/mobilenetwork/resource-manager/Microsoft.MobileNetwork/stable/2023-06-01/examples/SliceCreate.json
     */
    /**
     * Sample code: Create network slice.
     *
     * @param manager Entry point to MobileNetworkManager.
     */
    public static void createNetworkSlice(com.azure.resourcemanager.mobilenetwork.MobileNetworkManager manager) {
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
