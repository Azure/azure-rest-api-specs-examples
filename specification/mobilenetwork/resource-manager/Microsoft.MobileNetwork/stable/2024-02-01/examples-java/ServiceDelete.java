
/**
 * Samples for Services Delete.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/mobilenetwork/resource-manager/Microsoft.MobileNetwork/stable/2024-02-01/examples/ServiceDelete.
     * json
     */
    /**
     * Sample code: Delete service.
     * 
     * @param manager Entry point to MobileNetworkManager.
     */
    public static void deleteService(com.azure.resourcemanager.mobilenetwork.MobileNetworkManager manager) {
        manager.services().delete("rg1", "testMobileNetwork", "TestService", com.azure.core.util.Context.NONE);
    }
}
