
import com.azure.resourcemanager.search.models.CheckNameAvailabilityInput;

/**
 * Samples for Services CheckNameAvailability.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-01-preview/SearchCheckNameAvailability.json
     */
    /**
     * Sample code: SearchCheckNameAvailability.
     * 
     * @param manager Entry point to SearchServiceManager.
     */
    public static void searchCheckNameAvailability(com.azure.resourcemanager.search.SearchServiceManager manager) {
        manager.serviceClient().getServices().checkNameAvailabilityWithResponse(
            new CheckNameAvailabilityInput().withName("mysearchservice"), com.azure.core.util.Context.NONE);
    }
}
