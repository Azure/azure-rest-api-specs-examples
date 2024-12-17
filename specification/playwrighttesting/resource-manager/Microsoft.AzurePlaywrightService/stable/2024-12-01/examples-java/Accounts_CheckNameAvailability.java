
import com.azure.resourcemanager.playwrighttesting.models.CheckNameAvailabilityRequest;

/**
 * Samples for Accounts CheckNameAvailability.
 */
public final class Main {
    /*
     * x-ms-original-file: 2024-12-01/Accounts_CheckNameAvailability.json
     */
    /**
     * Sample code: Accounts_CheckNameAvailability.
     * 
     * @param manager Entry point to PlaywrightTestingManager.
     */
    public static void
        accountsCheckNameAvailability(com.azure.resourcemanager.playwrighttesting.PlaywrightTestingManager manager) {
        manager.accounts().checkNameAvailabilityWithResponse(new CheckNameAvailabilityRequest().withName("dummyName")
            .withType("Microsoft.AzurePlaywrightService/Accounts"), com.azure.core.util.Context.NONE);
    }
}
