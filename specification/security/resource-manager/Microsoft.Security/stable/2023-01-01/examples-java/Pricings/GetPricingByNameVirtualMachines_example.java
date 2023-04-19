/** Samples for Pricings Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/security/resource-manager/Microsoft.Security/stable/2023-01-01/examples/Pricings/GetPricingByNameVirtualMachines_example.json
     */
    /**
     * Sample code: Get pricings on subscription - VirtualMachines plan.
     *
     * @param manager Entry point to SecurityManager.
     */
    public static void getPricingsOnSubscriptionVirtualMachinesPlan(
        com.azure.resourcemanager.security.SecurityManager manager) {
        manager.pricings().getWithResponse("VirtualMachines", com.azure.core.util.Context.NONE);
    }
}
