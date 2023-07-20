/** Samples for Services ListByMobileNetwork. */
public final class Main {
    /*
     * x-ms-original-file: specification/mobilenetwork/resource-manager/Microsoft.MobileNetwork/stable/2023-06-01/examples/ServiceListByMobileNetwork.json
     */
    /**
     * Sample code: List services in a mobile network.
     *
     * @param manager Entry point to MobileNetworkManager.
     */
    public static void listServicesInAMobileNetwork(
        com.azure.resourcemanager.mobilenetwork.MobileNetworkManager manager) {
        manager
            .services()
            .listByMobileNetwork("testResourceGroupName", "testMobileNetwork", com.azure.core.util.Context.NONE);
    }
}
