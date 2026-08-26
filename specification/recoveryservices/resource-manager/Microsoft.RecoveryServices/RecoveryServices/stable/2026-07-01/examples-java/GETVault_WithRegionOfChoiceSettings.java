
/**
 * Samples for Vaults GetByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-07-01/GETVault_WithRegionOfChoiceSettings.json
     */
    /**
     * Sample code: Get Recovery Services Vault With Region Of Choice Settings.
     * 
     * @param manager Entry point to RecoveryServicesManager.
     */
    public static void getRecoveryServicesVaultWithRegionOfChoiceSettings(
        com.azure.resourcemanager.recoveryservices.RecoveryServicesManager manager) {
        manager.vaults().getByResourceGroupWithResponse("Default-RecoveryServices-ResourceGroup", "swaggerExample",
            com.azure.core.util.Context.NONE);
    }
}
