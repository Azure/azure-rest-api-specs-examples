
import com.azure.resourcemanager.recoveryservices.fluent.models.VaultExtendedInfoResourceInner;

/**
 * Samples for VaultExtendedInfo CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-05-01/UpdateVaultExtendedInfo_Put.json
     */
    /**
     * Sample code: Put ExtendedInfo of Resource.
     * 
     * @param manager Entry point to RecoveryServicesManager.
     */
    public static void
        putExtendedInfoOfResource(com.azure.resourcemanager.recoveryservices.RecoveryServicesManager manager) {
        manager.vaultExtendedInfoes().createOrUpdateWithResponse("Default-RecoveryServices-ResourceGroup",
            "swaggerExample",
            new VaultExtendedInfoResourceInner().withIntegrityKey("fakeTokenPlaceholder").withAlgorithm("None"),
            com.azure.core.util.Context.NONE);
    }
}
