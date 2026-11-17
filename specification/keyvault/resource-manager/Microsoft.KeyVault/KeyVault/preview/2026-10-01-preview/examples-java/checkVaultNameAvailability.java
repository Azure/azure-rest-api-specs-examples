
import com.azure.resourcemanager.keyvault.models.VaultCheckNameAvailabilityParameters;

/**
 * Samples for Vaults CheckNameAvailability.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-10-01-preview/checkVaultNameAvailability.json
     */
    /**
     * Sample code: Validate a vault name.
     * 
     * @param manager Entry point to KeyVaultManager.
     */
    public static void validateAVaultName(com.azure.resourcemanager.keyvault.KeyVaultManager manager) {
        manager.serviceClient().getVaults().checkNameAvailabilityWithResponse(
            new VaultCheckNameAvailabilityParameters().withName("sample-vault"), com.azure.core.util.Context.NONE);
    }
}
