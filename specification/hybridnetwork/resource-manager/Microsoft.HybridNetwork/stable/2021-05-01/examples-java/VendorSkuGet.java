import com.azure.core.util.Context;

/** Samples for VendorSkus Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/hybridnetwork/resource-manager/Microsoft.HybridNetwork/stable/2021-05-01/examples/VendorSkuGet.json
     */
    /**
     * Sample code: Get the sku of vendor resource.
     *
     * @param manager Entry point to HybridNetworkManager.
     */
    public static void getTheSkuOfVendorResource(com.azure.resourcemanager.hybridnetwork.HybridNetworkManager manager) {
        manager.vendorSkus().getWithResponse("TestVendor", "TestSku", Context.NONE);
    }
}
