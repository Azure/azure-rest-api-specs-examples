import com.azure.core.util.Context;

/** Samples for Vendors Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/hybridnetwork/resource-manager/Microsoft.HybridNetwork/stable/2021-05-01/examples/VendorGet.json
     */
    /**
     * Sample code: Get Vendor resource.
     *
     * @param manager Entry point to HybridNetworkManager.
     */
    public static void getVendorResource(com.azure.resourcemanager.hybridnetwork.HybridNetworkManager manager) {
        manager.vendors().getWithResponse("TestVendor", Context.NONE);
    }
}
