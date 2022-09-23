import com.azure.core.util.Context;

/** Samples for Pricings Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/security/resource-manager/Microsoft.Security/stable/2022-03-01/examples/Pricings/GetPricingByName_example.json
     */
    /**
     * Sample code: Get pricings on subscription.
     *
     * @param manager Entry point to SecurityManager.
     */
    public static void getPricingsOnSubscription(com.azure.resourcemanager.security.SecurityManager manager) {
        manager.pricings().getWithResponse("VirtualMachines", Context.NONE);
    }
}
