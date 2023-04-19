/** Samples for Pricings Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/security/resource-manager/Microsoft.Security/stable/2023-01-01/examples/Pricings/GetPricingByNameDns_example.json
     */
    /**
     * Sample code: Get pricings on subscription - Dns plan.
     *
     * @param manager Entry point to SecurityManager.
     */
    public static void getPricingsOnSubscriptionDnsPlan(com.azure.resourcemanager.security.SecurityManager manager) {
        manager.pricings().getWithResponse("Dns", com.azure.core.util.Context.NONE);
    }
}
