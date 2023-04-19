/** Samples for Pricings Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/security/resource-manager/Microsoft.Security/stable/2023-01-01/examples/Pricings/GetPricingByNameStorageAccounts_example.json
     */
    /**
     * Sample code: Get pricings on subscription - StorageAccounts plan.
     *
     * @param manager Entry point to SecurityManager.
     */
    public static void getPricingsOnSubscriptionStorageAccountsPlan(
        com.azure.resourcemanager.security.SecurityManager manager) {
        manager.pricings().getWithResponse("StorageAccounts", com.azure.core.util.Context.NONE);
    }
}
