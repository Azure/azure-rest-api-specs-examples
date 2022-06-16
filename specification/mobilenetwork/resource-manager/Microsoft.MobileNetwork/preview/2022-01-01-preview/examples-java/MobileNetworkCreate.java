import com.azure.resourcemanager.mobilenetwork.models.PlmnId;

/** Samples for MobileNetworks CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/mobilenetwork/resource-manager/Microsoft.MobileNetwork/preview/2022-01-01-preview/examples/MobileNetworkCreate.json
     */
    /**
     * Sample code: Create mobile network.
     *
     * @param manager Entry point to MobileNetworkManager.
     */
    public static void createMobileNetwork(com.azure.resourcemanager.mobilenetwork.MobileNetworkManager manager) {
        manager
            .mobileNetworks()
            .define("testMobileNetwork")
            .withRegion("eastus")
            .withExistingResourceGroup("rg1")
            .withPublicLandMobileNetworkIdentifier(new PlmnId().withMcc("001").withMnc("01"))
            .create();
    }
}
