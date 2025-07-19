
/**
 * Samples for Operations List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-03-31/OperationsList.json
     */
    /**
     * Sample code: Get a list of Payment HSM operations.
     * 
     * @param manager Entry point to HardwareSecurityModulesManager.
     */
    public static void getAListOfPaymentHSMOperations(
        com.azure.resourcemanager.hardwaresecuritymodules.HardwareSecurityModulesManager manager) {
        manager.operations().list(com.azure.core.util.Context.NONE);
    }
}
