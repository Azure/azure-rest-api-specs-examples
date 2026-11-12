
import com.azure.resourcemanager.sovereign.models.CheckNameAvailabilityRequest;

/**
 * Samples for Operations CheckNameAvailabilityLocal.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-10-31-preview/Operations_CheckNameAvailabilityLocal.json
     */
    /**
     * Sample code: Operations_CheckNameAvailabilityLocal.
     * 
     * @param manager Entry point to SovereignManager.
     */
    public static void
        operationsCheckNameAvailabilityLocal(com.azure.resourcemanager.sovereign.SovereignManager manager) {
        manager.operations().checkNameAvailabilityLocalWithResponse("northeurope",
            new CheckNameAvailabilityRequest().withName("name1").withType("Microsoft.Sovereign/sovereignViews"),
            com.azure.core.util.Context.NONE);
    }
}
