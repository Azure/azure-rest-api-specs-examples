/** Samples for Vendors CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/hybridnetwork/resource-manager/Microsoft.HybridNetwork/stable/2021-05-01/examples/VendorCreate.json
     */
    /**
     * Sample code: Create or update Vendor resource.
     *
     * @param manager Entry point to HybridNetworkManager.
     */
    public static void createOrUpdateVendorResource(
        com.azure.resourcemanager.hybridnetwork.HybridNetworkManager manager) {
        manager.vendors().define("TestVendor").create();
    }
}
