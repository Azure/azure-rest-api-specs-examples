import com.azure.core.util.Context;

/** Samples for VendorSkus List. */
public final class Main {
    /*
     * x-ms-original-file: specification/hybridnetwork/resource-manager/Microsoft.HybridNetwork/stable/2021-05-01/examples/VendorSkuListByVendor.json
     */
    /**
     * Sample code: List all the vendor skus of vendor resource.
     *
     * @param manager Entry point to HybridNetworkManager.
     */
    public static void listAllTheVendorSkusOfVendorResource(
        com.azure.resourcemanager.hybridnetwork.HybridNetworkManager manager) {
        manager.vendorSkus().list("TestVendor", Context.NONE);
    }
}
