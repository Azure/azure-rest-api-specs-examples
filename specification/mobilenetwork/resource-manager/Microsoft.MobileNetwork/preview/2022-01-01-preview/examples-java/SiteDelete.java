import com.azure.core.util.Context;

/** Samples for Sites Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/mobilenetwork/resource-manager/Microsoft.MobileNetwork/preview/2022-01-01-preview/examples/SiteDelete.json
     */
    /**
     * Sample code: Delete mobile network site.
     *
     * @param manager Entry point to MobileNetworkManager.
     */
    public static void deleteMobileNetworkSite(com.azure.resourcemanager.mobilenetwork.MobileNetworkManager manager) {
        manager.sites().delete("rg1", "testMobileNetwork", "testSite", Context.NONE);
    }
}
