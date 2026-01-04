
import com.azure.resourcemanager.keyvault.models.CheckMhsmNameAvailabilityParameters;

/**
 * Samples for ManagedHsms CheckMhsmNameAvailability.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-05-01/ManagedHsm_checkMhsmNameAvailability.json
     */
    /**
     * Sample code: Validate a managed hsm name.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void validateAManagedHsmName(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.vaults().manager().serviceClient().getManagedHsms().checkMhsmNameAvailabilityWithResponse(
            new CheckMhsmNameAvailabilityParameters().withName("sample-mhsm"), com.azure.core.util.Context.NONE);
    }
}
