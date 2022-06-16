/** Samples for DataNetworks CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/mobilenetwork/resource-manager/Microsoft.MobileNetwork/preview/2022-03-01-preview/examples/DataNetworkCreate.json
     */
    /**
     * Sample code: Create mobile network dataNetwork.
     *
     * @param manager Entry point to MobileNetworkManager.
     */
    public static void createMobileNetworkDataNetwork(
        com.azure.resourcemanager.mobilenetwork.MobileNetworkManager manager) {
        manager
            .dataNetworks()
            .define("testDataNetwork")
            .withRegion("eastus")
            .withExistingMobileNetwork("rg1", "testMobileNetwork")
            .withDescription("myFavouriteDataNetwork")
            .create();
    }
}
