
/**
 * Samples for MdeOnboardings List.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/security/resource-manager/Microsoft.Security/preview/2021-10-01-preview/examples/MdeOnboardings/
     * ListMdeOnboardings_example.json
     */
    /**
     * Sample code: The configuration or data needed to onboard the machine to MDE.
     * 
     * @param manager Entry point to SecurityManager.
     */
    public static void theConfigurationOrDataNeededToOnboardTheMachineToMDE(
        com.azure.resourcemanager.security.SecurityManager manager) {
        manager.mdeOnboardings().listWithResponse(com.azure.core.util.Context.NONE);
    }
}
