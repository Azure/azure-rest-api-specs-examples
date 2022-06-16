import com.azure.core.util.Context;

/** Samples for Sites ListByMobileNetwork. */
public final class Main {
    /*
     * x-ms-original-file: specification/mobilenetwork/resource-manager/Microsoft.MobileNetwork/preview/2022-01-01-preview/examples/SiteListByMobileNetwork.json
     */
    /**
     * Sample code: List mobile network sites in a mobile network.
     *
     * @param manager Entry point to MobileNetworkManager.
     */
    public static void listMobileNetworkSitesInAMobileNetwork(
        com.azure.resourcemanager.mobilenetwork.MobileNetworkManager manager) {
        manager.sites().listByMobileNetwork("rg1", "testMobileNetwork", Context.NONE);
    }
}
