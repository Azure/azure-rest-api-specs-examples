
/**
 * Samples for DedicatedHsm GetByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/hardwaresecuritymodules/resource-manager/Microsoft.HardwareSecurityModules/preview/2024-06-30-
     * preview/examples/DedicatedHsm_Get.json
     */
    /**
     * Sample code: Get a dedicated HSM.
     * 
     * @param manager Entry point to HardwareSecurityModulesManager.
     */
    public static void
        getADedicatedHSM(com.azure.resourcemanager.hardwaresecuritymodules.HardwareSecurityModulesManager manager) {
        manager.dedicatedHsms().getByResourceGroupWithResponse("hsm-group", "hsm1", com.azure.core.util.Context.NONE);
    }
}
