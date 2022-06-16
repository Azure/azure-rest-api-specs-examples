import com.azure.core.util.Context;

/** Samples for MobileNetworks GetByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/mobilenetwork/resource-manager/Microsoft.MobileNetwork/preview/2022-01-01-preview/examples/MobileNetworkGet.json
     */
    /**
     * Sample code: Get mobile network.
     *
     * @param manager Entry point to MobileNetworkManager.
     */
    public static void getMobileNetwork(com.azure.resourcemanager.mobilenetwork.MobileNetworkManager manager) {
        manager.mobileNetworks().getByResourceGroupWithResponse("rg1", "testMobileNetwork", Context.NONE);
    }
}
