
/**
 * Samples for DedicatedHsm ListByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-03-31/DedicatedHsm_ListByResourceGroup.json
     */
    /**
     * Sample code: List dedicated HSM devices in a resource group.
     * 
     * @param manager Entry point to HardwareSecurityModulesManager.
     */
    public static void listDedicatedHSMDevicesInAResourceGroup(
        com.azure.resourcemanager.hardwaresecuritymodules.HardwareSecurityModulesManager manager) {
        manager.dedicatedHsms().listByResourceGroup("hsm-group", null, com.azure.core.util.Context.NONE);
    }
}
