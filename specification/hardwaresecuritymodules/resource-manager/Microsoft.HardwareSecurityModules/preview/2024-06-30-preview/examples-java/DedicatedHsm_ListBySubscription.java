
/**
 * Samples for DedicatedHsm List.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/hardwaresecuritymodules/resource-manager/Microsoft.HardwareSecurityModules/preview/2024-06-30-
     * preview/examples/DedicatedHsm_ListBySubscription.json
     */
    /**
     * Sample code: List dedicated HSM devices in a subscription.
     * 
     * @param manager Entry point to HardwareSecurityModulesManager.
     */
    public static void listDedicatedHSMDevicesInASubscription(
        com.azure.resourcemanager.hardwaresecuritymodules.HardwareSecurityModulesManager manager) {
        manager.dedicatedHsms().list(null, com.azure.core.util.Context.NONE);
    }
}
