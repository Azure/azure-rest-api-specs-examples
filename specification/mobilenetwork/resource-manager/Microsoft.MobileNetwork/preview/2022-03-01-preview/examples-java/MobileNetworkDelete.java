import com.azure.core.util.Context;

/** Samples for MobileNetworks Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/mobilenetwork/resource-manager/Microsoft.MobileNetwork/preview/2022-03-01-preview/examples/MobileNetworkDelete.json
     */
    /**
     * Sample code: Delete mobile network.
     *
     * @param manager Entry point to MobileNetworkManager.
     */
    public static void deleteMobileNetwork(com.azure.resourcemanager.mobilenetwork.MobileNetworkManager manager) {
        manager.mobileNetworks().delete("rg1", "testMobileNetwork", Context.NONE);
    }
}
