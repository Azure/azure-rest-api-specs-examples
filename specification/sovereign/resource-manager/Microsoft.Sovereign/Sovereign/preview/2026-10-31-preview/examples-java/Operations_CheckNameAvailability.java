
import com.azure.resourcemanager.sovereign.models.CheckNameAvailabilityRequest;

/**
 * Samples for Operations CheckNameAvailability.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-10-31-preview/Operations_CheckNameAvailability.json
     */
    /**
     * Sample code: Operations_CheckNameAvailability.
     * 
     * @param manager Entry point to SovereignManager.
     */
    public static void operationsCheckNameAvailability(com.azure.resourcemanager.sovereign.SovereignManager manager) {
        manager.operations().checkNameAvailabilityWithResponse(
            new CheckNameAvailabilityRequest().withName("name1").withType("Microsoft.Sovereign/sovereignViews"),
            com.azure.core.util.Context.NONE);
    }
}
