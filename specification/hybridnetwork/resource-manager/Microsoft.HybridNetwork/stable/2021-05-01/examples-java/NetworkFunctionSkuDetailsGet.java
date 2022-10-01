import com.azure.core.util.Context;

/** Samples for NetworkFunctionVendorSkus ListBySku. */
public final class Main {
    /*
     * x-ms-original-file: specification/hybridnetwork/resource-manager/Microsoft.HybridNetwork/stable/2021-05-01/examples/NetworkFunctionSkuDetailsGet.json
     */
    /**
     * Sample code: Get network function sku details.
     *
     * @param manager Entry point to HybridNetworkManager.
     */
    public static void getNetworkFunctionSkuDetails(
        com.azure.resourcemanager.hybridnetwork.HybridNetworkManager manager) {
        manager.networkFunctionVendorSkus().listBySku("testVendor", "testSku", Context.NONE);
    }
}
