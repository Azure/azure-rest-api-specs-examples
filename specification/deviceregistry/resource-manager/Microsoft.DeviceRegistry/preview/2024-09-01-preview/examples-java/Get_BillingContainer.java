
/**
 * Samples for BillingContainers Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2024-09-01-preview/Get_BillingContainer.json
     */
    /**
     * Sample code: Get_BillingContainer.
     * 
     * @param manager Entry point to DeviceRegistryManager.
     */
    public static void getBillingContainer(com.azure.resourcemanager.deviceregistry.DeviceRegistryManager manager) {
        manager.billingContainers().getWithResponse("my-billingContainer", com.azure.core.util.Context.NONE);
    }
}
