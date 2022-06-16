import com.azure.core.util.Context;

/** Samples for Services ListByMobileNetwork. */
public final class Main {
    /*
     * x-ms-original-file: specification/mobilenetwork/resource-manager/Microsoft.MobileNetwork/preview/2022-03-01-preview/examples/ServiceListByMobileNetwork.json
     */
    /**
     * Sample code: List services in a mobile network.
     *
     * @param manager Entry point to MobileNetworkManager.
     */
    public static void listServicesInAMobileNetwork(
        com.azure.resourcemanager.mobilenetwork.MobileNetworkManager manager) {
        manager.services().listByMobileNetwork("testResourceGroupName", "testMobileNetwork", Context.NONE);
    }
}
