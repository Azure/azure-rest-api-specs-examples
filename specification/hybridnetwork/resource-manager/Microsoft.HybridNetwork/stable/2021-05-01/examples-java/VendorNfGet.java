import com.azure.core.util.Context;

/** Samples for VendorNetworkFunctions Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/hybridnetwork/resource-manager/Microsoft.HybridNetwork/stable/2021-05-01/examples/VendorNfGet.json
     */
    /**
     * Sample code: Get vendor network function sub resource by service key of network function.
     *
     * @param manager Entry point to HybridNetworkManager.
     */
    public static void getVendorNetworkFunctionSubResourceByServiceKeyOfNetworkFunction(
        com.azure.resourcemanager.hybridnetwork.HybridNetworkManager manager) {
        manager.vendorNetworkFunctions().getWithResponse("eastus", "testVendor", "testServiceKey", Context.NONE);
    }
}
