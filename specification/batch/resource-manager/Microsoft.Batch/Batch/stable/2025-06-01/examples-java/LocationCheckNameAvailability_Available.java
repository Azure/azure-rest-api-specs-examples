
import com.azure.resourcemanager.batch.models.CheckNameAvailabilityParameters;
import com.azure.resourcemanager.batch.models.ResourceType;

/**
 * Samples for Location CheckNameAvailability.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-06-01/LocationCheckNameAvailability_Available.json
     */
    /**
     * Sample code: LocationCheckNameAvailability_Available.
     * 
     * @param manager Entry point to BatchManager.
     */
    public static void locationCheckNameAvailabilityAvailable(com.azure.resourcemanager.batch.BatchManager manager) {
        manager.locations()
            .checkNameAvailabilityWithResponse("japaneast", new CheckNameAvailabilityParameters()
                .withName("newaccountname").withType(ResourceType.MICROSOFT_BATCH_BATCH_ACCOUNTS),
                com.azure.core.util.Context.NONE);
    }
}
