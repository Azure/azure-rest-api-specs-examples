
/**
 * Samples for DedicatedHsm List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-12-01-preview/PaymentHsm_ListBySubscription.json
     */
    /**
     * Sample code: List dedicated HSM devices in a subscription including payment HSM.
     * 
     * @param manager Entry point to HardwareSecurityModulesManager.
     */
    public static void listDedicatedHSMDevicesInASubscriptionIncludingPaymentHSM(
        com.azure.resourcemanager.hardwaresecuritymodules.HardwareSecurityModulesManager manager) {
        manager.dedicatedHsms().list(null, com.azure.core.util.Context.NONE);
    }
}
