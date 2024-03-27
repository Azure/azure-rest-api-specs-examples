
/**
 * Samples for Products CountDevices.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/sphere/resource-manager/Microsoft.AzureSphere/stable/2024-04-01/examples/PostCountDevicesProduct.
     * json
     */
    /**
     * Sample code: Products_CountDevices.
     * 
     * @param manager Entry point to AzureSphereManager.
     */
    public static void productsCountDevices(com.azure.resourcemanager.sphere.AzureSphereManager manager) {
        manager.products().countDevicesWithResponse("MyResourceGroup1", "MyCatalog1", "MyProduct1",
            com.azure.core.util.Context.NONE);
    }
}
