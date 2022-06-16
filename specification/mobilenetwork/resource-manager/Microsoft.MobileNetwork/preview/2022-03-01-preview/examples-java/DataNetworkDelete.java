import com.azure.core.util.Context;

/** Samples for DataNetworks Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/mobilenetwork/resource-manager/Microsoft.MobileNetwork/preview/2022-03-01-preview/examples/DataNetworkDelete.json
     */
    /**
     * Sample code: Delete mobile network dataNetwork.
     *
     * @param manager Entry point to MobileNetworkManager.
     */
    public static void deleteMobileNetworkDataNetwork(
        com.azure.resourcemanager.mobilenetwork.MobileNetworkManager manager) {
        manager.dataNetworks().delete("rg1", "testMobileNetwork", "testDataNetwork", Context.NONE);
    }
}
