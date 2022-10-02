import com.azure.core.util.Context;

/** Samples for VendorSkuPreview Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/hybridnetwork/resource-manager/Microsoft.HybridNetwork/stable/2021-05-01/examples/VendorSkuPreviewDelete.json
     */
    /**
     * Sample code: Delete preview subscription of vendor sku sub resource.
     *
     * @param manager Entry point to HybridNetworkManager.
     */
    public static void deletePreviewSubscriptionOfVendorSkuSubResource(
        com.azure.resourcemanager.hybridnetwork.HybridNetworkManager manager) {
        manager.vendorSkuPreviews().delete("TestVendor", "TestSku", "previewSub", Context.NONE);
    }
}
