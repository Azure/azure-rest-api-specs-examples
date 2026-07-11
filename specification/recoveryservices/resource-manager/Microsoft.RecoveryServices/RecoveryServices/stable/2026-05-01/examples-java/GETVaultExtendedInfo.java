
/**
 * Samples for VaultExtendedInfo Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-05-01/GETVaultExtendedInfo.json
     */
    /**
     * Sample code: Get ExtendedInfo of Resource.
     * 
     * @param manager Entry point to RecoveryServicesManager.
     */
    public static void
        getExtendedInfoOfResource(com.azure.resourcemanager.recoveryservices.RecoveryServicesManager manager) {
        manager.vaultExtendedInfoes().getWithResponse("Default-RecoveryServices-ResourceGroup", "swaggerExample",
            com.azure.core.util.Context.NONE);
    }
}
