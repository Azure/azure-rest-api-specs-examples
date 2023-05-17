/** Samples for VaultExtendedInfo Update. */
public final class Main {
    /*
     * x-ms-original-file: specification/recoveryservices/resource-manager/Microsoft.RecoveryServices/stable/2023-02-01/examples/UpdateVaultExtendedInfo.json
     */
    /**
     * Sample code: PATCH ExtendedInfo of Resource.
     *
     * @param manager Entry point to RecoveryServicesManager.
     */
    public static void pATCHExtendedInfoOfResource(
        com.azure.resourcemanager.recoveryservices.RecoveryServicesManager manager) {
        manager
            .vaultExtendedInfoes()
            .updateWithResponse(
                "Default-RecoveryServices-ResourceGroup", "swaggerExample", null, com.azure.core.util.Context.NONE);
    }
}
