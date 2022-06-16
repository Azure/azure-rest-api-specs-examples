import com.azure.core.util.Context;

/** Samples for Sites Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/mobilenetwork/resource-manager/Microsoft.MobileNetwork/preview/2022-01-01-preview/examples/SiteGet.json
     */
    /**
     * Sample code: Get mobile network site.
     *
     * @param manager Entry point to MobileNetworkManager.
     */
    public static void getMobileNetworkSite(com.azure.resourcemanager.mobilenetwork.MobileNetworkManager manager) {
        manager.sites().getWithResponse("rg1", "testMobileNetwork", "testSite", Context.NONE);
    }
}
