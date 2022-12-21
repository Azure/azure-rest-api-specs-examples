import com.azure.core.util.Context;

/** Samples for DataNetworks Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/mobilenetwork/resource-manager/Microsoft.MobileNetwork/stable/2022-11-01/examples/DataNetworkGet.json
     */
    /**
     * Sample code: Get data network.
     *
     * @param manager Entry point to MobileNetworkManager.
     */
    public static void getDataNetwork(com.azure.resourcemanager.mobilenetwork.MobileNetworkManager manager) {
        manager.dataNetworks().getWithResponse("rg1", "testMobileNetwork", "testDataNetwork", Context.NONE);
    }
}
