
/**
 * Samples for VaultExtendedInfo CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/recoveryservices/resource-manager/Microsoft.RecoveryServices/stable/2024-04-01/examples/
     * UpdateVaultExtendedInfo.json
     */
    /**
     * Sample code: Put ExtendedInfo of Resource.
     * 
     * @param manager Entry point to RecoveryServicesManager.
     */
    public static void
        putExtendedInfoOfResource(com.azure.resourcemanager.recoveryservices.RecoveryServicesManager manager) {
        manager.vaultExtendedInfoes().createOrUpdateWithResponse("Default-RecoveryServices-ResourceGroup",
            "swaggerExample", null, com.azure.core.util.Context.NONE);
    }
}
