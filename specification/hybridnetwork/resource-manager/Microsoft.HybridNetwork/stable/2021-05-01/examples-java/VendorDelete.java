import com.azure.core.util.Context;

/** Samples for Vendors Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/hybridnetwork/resource-manager/Microsoft.HybridNetwork/stable/2021-05-01/examples/VendorDelete.json
     */
    /**
     * Sample code: Delete vendor resource.
     *
     * @param manager Entry point to HybridNetworkManager.
     */
    public static void deleteVendorResource(com.azure.resourcemanager.hybridnetwork.HybridNetworkManager manager) {
        manager.vendors().delete("TestVendor", Context.NONE);
    }
}
