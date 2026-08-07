
/**
 * Samples for DedicatedHsm GetByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-12-01-preview/PaymentHsm_Get.json
     */
    /**
     * Sample code: Get a payment HSM.
     * 
     * @param manager Entry point to HardwareSecurityModulesManager.
     */
    public static void
        getAPaymentHSM(com.azure.resourcemanager.hardwaresecuritymodules.HardwareSecurityModulesManager manager) {
        manager.dedicatedHsms().getByResourceGroupWithResponse("hsm-group", "hsm1", com.azure.core.util.Context.NONE);
    }
}
