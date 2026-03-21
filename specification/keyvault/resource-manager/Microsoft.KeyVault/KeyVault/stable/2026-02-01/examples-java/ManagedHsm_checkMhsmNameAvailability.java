
import com.azure.resourcemanager.keyvault.models.CheckMhsmNameAvailabilityParameters;

/**
 * Samples for ManagedHsms CheckMhsmNameAvailability.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-02-01/ManagedHsm_checkMhsmNameAvailability.json
     */
    /**
     * Sample code: Validate a managed hsm name.
     * 
     * @param manager Entry point to KeyVaultManager.
     */
    public static void validateAManagedHsmName(com.azure.resourcemanager.keyvault.KeyVaultManager manager) {
        manager.serviceClient().getManagedHsms().checkMhsmNameAvailabilityWithResponse(
            new CheckMhsmNameAvailabilityParameters().withName("sample-mhsm"), com.azure.core.util.Context.NONE);
    }
}
