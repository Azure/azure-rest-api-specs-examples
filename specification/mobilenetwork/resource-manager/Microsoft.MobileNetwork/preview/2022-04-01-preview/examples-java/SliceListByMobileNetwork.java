import com.azure.core.util.Context;

/** Samples for Slices ListByMobileNetwork. */
public final class Main {
    /*
     * x-ms-original-file: specification/mobilenetwork/resource-manager/Microsoft.MobileNetwork/preview/2022-04-01-preview/examples/SliceListByMobileNetwork.json
     */
    /**
     * Sample code: List network slices in a mobile network.
     *
     * @param manager Entry point to MobileNetworkManager.
     */
    public static void listNetworkSlicesInAMobileNetwork(
        com.azure.resourcemanager.mobilenetwork.MobileNetworkManager manager) {
        manager.slices().listByMobileNetwork("rg1", "testMobileNetwork", Context.NONE);
    }
}
