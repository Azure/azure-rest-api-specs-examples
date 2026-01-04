
import com.azure.resourcemanager.keyvault.models.VaultCheckNameAvailabilityParameters;

/**
 * Samples for Vaults CheckNameAvailability.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-05-01/checkVaultNameAvailability.json
     */
    /**
     * Sample code: Validate a vault name.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void validateAVaultName(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.vaults().manager().serviceClient().getVaults().checkNameAvailabilityWithResponse(
            new VaultCheckNameAvailabilityParameters().withName("sample-vault"), com.azure.core.util.Context.NONE);
    }
}
