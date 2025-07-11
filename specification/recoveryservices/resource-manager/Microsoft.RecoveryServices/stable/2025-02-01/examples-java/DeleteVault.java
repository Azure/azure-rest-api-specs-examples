
/**
 * Samples for Vaults Delete.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/recoveryservices/resource-manager/Microsoft.RecoveryServices/stable/2025-02-01/examples/DeleteVault
     * .json
     */
    /**
     * Sample code: Delete Recovery Services Vault.
     * 
     * @param manager Entry point to RecoveryServicesManager.
     */
    public static void
        deleteRecoveryServicesVault(com.azure.resourcemanager.recoveryservices.RecoveryServicesManager manager) {
        manager.vaults().delete("Default-RecoveryServices-ResourceGroup", "swaggerExample",
            com.azure.core.util.Context.NONE);
    }
}
