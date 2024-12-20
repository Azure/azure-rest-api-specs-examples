
/**
 * Samples for CustomerSubscriptions Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/azurestack/resource-manager/Microsoft.AzureStack/stable/2022-06-01/examples/CustomerSubscription/
     * Get.json
     */
    /**
     * Sample code: Returns the specified product.
     * 
     * @param manager Entry point to AzureStackManager.
     */
    public static void returnsTheSpecifiedProduct(com.azure.resourcemanager.azurestack.AzureStackManager manager) {
        manager.customerSubscriptions().getWithResponse("azurestack", "testregistration",
            "E09A4E93-29A7-4EBA-A6D4-76202383F07F", com.azure.core.util.Context.NONE);
    }
}
