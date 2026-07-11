
import com.azure.resourcemanager.recoveryservices.fluent.models.VaultExtendedInfoResourceInner;

/**
 * Samples for VaultExtendedInfo Update.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-05-01/UpdateVaultExtendedInfo.json
     */
    /**
     * Sample code: PATCH ExtendedInfo of Resource.
     * 
     * @param manager Entry point to RecoveryServicesManager.
     */
    public static void
        pATCHExtendedInfoOfResource(com.azure.resourcemanager.recoveryservices.RecoveryServicesManager manager) {
        manager.vaultExtendedInfoes().updateWithResponse("Default-RecoveryServices-ResourceGroup", "swaggerExample",
            new VaultExtendedInfoResourceInner().withIntegrityKey("fakeTokenPlaceholder").withAlgorithm("None"),
            com.azure.core.util.Context.NONE);
    }
}
