import com.azure.core.util.Context;

/** Samples for VendorSkuPreview List. */
public final class Main {
    /*
     * x-ms-original-file: specification/hybridnetwork/resource-manager/Microsoft.HybridNetwork/stable/2021-05-01/examples/VendorSkuPreviewListBySku.json
     */
    /**
     * Sample code: List all preview subscriptions of vendor sku sub resource.
     *
     * @param manager Entry point to HybridNetworkManager.
     */
    public static void listAllPreviewSubscriptionsOfVendorSkuSubResource(
        com.azure.resourcemanager.hybridnetwork.HybridNetworkManager manager) {
        manager.vendorSkuPreviews().list("TestVendor", "TestSku", Context.NONE);
    }
}
