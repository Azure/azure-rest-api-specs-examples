/** Samples for VendorSkuPreview CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/hybridnetwork/resource-manager/Microsoft.HybridNetwork/stable/2021-05-01/examples/VendorSkuPreviewCreate.json
     */
    /**
     * Sample code: Create or update preview subscription of vendor sku sub resource.
     *
     * @param manager Entry point to HybridNetworkManager.
     */
    public static void createOrUpdatePreviewSubscriptionOfVendorSkuSubResource(
        com.azure.resourcemanager.hybridnetwork.HybridNetworkManager manager) {
        manager.vendorSkuPreviews().define("previewSub").withExistingVendorSku("TestVendor", "TestSku").create();
    }
}
