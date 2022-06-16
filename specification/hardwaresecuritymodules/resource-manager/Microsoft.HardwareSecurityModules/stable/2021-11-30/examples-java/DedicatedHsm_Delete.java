import com.azure.core.util.Context;

/** Samples for DedicatedHsm Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/hardwaresecuritymodules/resource-manager/Microsoft.HardwareSecurityModules/stable/2021-11-30/examples/DedicatedHsm_Delete.json
     */
    /**
     * Sample code: Delete a dedicated HSM.
     *
     * @param manager Entry point to HardwareSecurityModulesManager.
     */
    public static void deleteADedicatedHSM(
        com.azure.resourcemanager.hardwaresecuritymodules.HardwareSecurityModulesManager manager) {
        manager.dedicatedHsms().delete("hsm-group", "hsm1", Context.NONE);
    }
}
