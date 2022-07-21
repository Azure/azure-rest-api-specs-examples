import com.azure.core.util.Context;

/** Samples for VaultExtendedInfo Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/recoveryservices/resource-manager/Microsoft.RecoveryServices/stable/2021-08-01/examples/GETVaultExtendedInfo.json
     */
    /**
     * Sample code: Get ExtendedInfo of Resource.
     *
     * @param manager Entry point to RecoveryServicesManager.
     */
    public static void getExtendedInfoOfResource(
        com.azure.resourcemanager.recoveryservices.RecoveryServicesManager manager) {
        manager
            .vaultExtendedInfoes()
            .getWithResponse("Default-RecoveryServices-ResourceGroup", "swaggerExample", Context.NONE);
    }
}
