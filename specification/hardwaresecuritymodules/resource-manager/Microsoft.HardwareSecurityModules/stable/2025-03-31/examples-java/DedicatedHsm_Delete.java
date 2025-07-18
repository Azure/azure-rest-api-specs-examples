
/**
 * Samples for DedicatedHsm Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-03-31/DedicatedHsm_Delete.json
     */
    /**
     * Sample code: Delete a dedicated HSM.
     * 
     * @param manager Entry point to HardwareSecurityModulesManager.
     */
    public static void
        deleteADedicatedHSM(com.azure.resourcemanager.hardwaresecuritymodules.HardwareSecurityModulesManager manager) {
        manager.dedicatedHsms().delete("hsm-group", "hsm1", com.azure.core.util.Context.NONE);
    }
}
