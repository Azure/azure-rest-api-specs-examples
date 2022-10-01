import com.azure.core.util.Context;

/** Samples for VendorSkuPreview Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/hybridnetwork/resource-manager/Microsoft.HybridNetwork/stable/2021-05-01/examples/VendorSkuPreviewGet.json
     */
    /**
     * Sample code: Get preview subscription of vendor sku sub resource.
     *
     * @param manager Entry point to HybridNetworkManager.
     */
    public static void getPreviewSubscriptionOfVendorSkuSubResource(
        com.azure.resourcemanager.hybridnetwork.HybridNetworkManager manager) {
        manager.vendorSkuPreviews().getWithResponse("TestVendor", "TestSku", "previewSub", Context.NONE);
    }
}
