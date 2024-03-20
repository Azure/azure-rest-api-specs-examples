
/**
 * Samples for DataNetworks Delete.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/mobilenetwork/resource-manager/Microsoft.MobileNetwork/stable/2024-02-01/examples/DataNetworkDelete
     * .json
     */
    /**
     * Sample code: Delete data network.
     * 
     * @param manager Entry point to MobileNetworkManager.
     */
    public static void deleteDataNetwork(com.azure.resourcemanager.mobilenetwork.MobileNetworkManager manager) {
        manager.dataNetworks().delete("rg1", "testMobileNetwork", "testDataNetwork", com.azure.core.util.Context.NONE);
    }
}
