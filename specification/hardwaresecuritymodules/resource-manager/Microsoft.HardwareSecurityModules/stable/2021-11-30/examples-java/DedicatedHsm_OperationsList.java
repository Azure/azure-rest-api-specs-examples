
/**
 * Samples for Operations List.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/hardwaresecuritymodules/resource-manager/Microsoft.HardwareSecurityModules/stable/2021-11-30/
     * examples/DedicatedHsm_OperationsList.json
     */
    /**
     * Sample code: Get a list of Dedicated HSM operations.
     * 
     * @param manager Entry point to HardwareSecurityModulesManager.
     */
    public static void getAListOfDedicatedHSMOperations(
        com.azure.resourcemanager.hardwaresecuritymodules.HardwareSecurityModulesManager manager) {
        manager.operations().list(com.azure.core.util.Context.NONE);
    }
}
