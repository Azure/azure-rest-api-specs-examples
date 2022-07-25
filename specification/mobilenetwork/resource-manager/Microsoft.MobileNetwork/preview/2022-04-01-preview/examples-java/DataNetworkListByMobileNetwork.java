import com.azure.core.util.Context;

/** Samples for DataNetworks ListByMobileNetwork. */
public final class Main {
    /*
     * x-ms-original-file: specification/mobilenetwork/resource-manager/Microsoft.MobileNetwork/preview/2022-04-01-preview/examples/DataNetworkListByMobileNetwork.json
     */
    /**
     * Sample code: List data networks in a mobile network.
     *
     * @param manager Entry point to MobileNetworkManager.
     */
    public static void listDataNetworksInAMobileNetwork(
        com.azure.resourcemanager.mobilenetwork.MobileNetworkManager manager) {
        manager.dataNetworks().listByMobileNetwork("rg1", "testMobileNetwork", Context.NONE);
    }
}
