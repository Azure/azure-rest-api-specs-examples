import com.azure.core.util.Context;

/** Samples for DataNetworks Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/mobilenetwork/resource-manager/Microsoft.MobileNetwork/preview/2022-01-01-preview/examples/DataNetworkGet.json
     */
    /**
     * Sample code: Get mobile network dataNetwork.
     *
     * @param manager Entry point to MobileNetworkManager.
     */
    public static void getMobileNetworkDataNetwork(
        com.azure.resourcemanager.mobilenetwork.MobileNetworkManager manager) {
        manager.dataNetworks().getWithResponse("rg1", "testMobileNetwork", "testDataNetwork", Context.NONE);
    }
}
