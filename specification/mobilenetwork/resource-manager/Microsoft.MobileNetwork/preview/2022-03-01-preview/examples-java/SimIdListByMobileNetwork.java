import com.azure.core.util.Context;

/** Samples for MobileNetworks ListSimIds. */
public final class Main {
    /*
     * x-ms-original-file: specification/mobilenetwork/resource-manager/Microsoft.MobileNetwork/preview/2022-03-01-preview/examples/SimIdListByMobileNetwork.json
     */
    /**
     * Sample code: List sim profile ids by network.
     *
     * @param manager Entry point to MobileNetworkManager.
     */
    public static void listSimProfileIdsByNetwork(
        com.azure.resourcemanager.mobilenetwork.MobileNetworkManager manager) {
        manager.mobileNetworks().listSimIds("rg", "testMobileNetworkName", Context.NONE);
    }
}
