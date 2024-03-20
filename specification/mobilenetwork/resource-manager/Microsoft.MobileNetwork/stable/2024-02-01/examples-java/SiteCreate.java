
/**
 * Samples for Sites CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/mobilenetwork/resource-manager/Microsoft.MobileNetwork/stable/2024-02-01/examples/SiteCreate.json
     */
    /**
     * Sample code: Create mobile network site.
     * 
     * @param manager Entry point to MobileNetworkManager.
     */
    public static void createMobileNetworkSite(com.azure.resourcemanager.mobilenetwork.MobileNetworkManager manager) {
        manager.sites().define("testSite").withRegion("testLocation")
            .withExistingMobileNetwork("rg1", "testMobileNetwork").create();
    }
}
