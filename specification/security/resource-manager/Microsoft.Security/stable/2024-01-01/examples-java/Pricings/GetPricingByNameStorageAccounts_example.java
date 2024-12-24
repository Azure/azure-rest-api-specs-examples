
/**
 * Samples for Pricings Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/security/resource-manager/Microsoft.Security/stable/2024-01-01/examples/Pricings/
     * GetPricingByNameStorageAccounts_example.json
     */
    /**
     * Sample code: Get pricings on subscription - StorageAccounts plan.
     * 
     * @param manager Entry point to SecurityManager.
     */
    public static void
        getPricingsOnSubscriptionStorageAccountsPlan(com.azure.resourcemanager.security.SecurityManager manager) {
        manager.pricings().getWithResponse("subscriptions/20ff7fc3-e762-44dd-bd96-b71116dcdc23", "StorageAccounts",
            com.azure.core.util.Context.NONE);
    }
}
