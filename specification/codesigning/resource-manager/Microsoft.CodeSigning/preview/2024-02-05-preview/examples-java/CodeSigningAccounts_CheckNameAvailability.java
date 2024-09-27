
import com.azure.resourcemanager.trustedsigning.models.CheckNameAvailability;

/**
 * Samples for CodeSigningAccounts CheckNameAvailability.
 */
public final class Main {
    /*
     * x-ms-original-file: 2024-02-05-preview/CodeSigningAccounts_CheckNameAvailability.json
     */
    /**
     * Sample code: Checks that the trusted signing account name is available.
     * 
     * @param manager Entry point to TrustedSigningManager.
     */
    public static void checksThatTheTrustedSigningAccountNameIsAvailable(
        com.azure.resourcemanager.trustedsigning.TrustedSigningManager manager) {
        manager.codeSigningAccounts().checkNameAvailabilityWithResponse(
            new CheckNameAvailability().withName("sample-account"), com.azure.core.util.Context.NONE);
    }
}
