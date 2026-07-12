
import com.azure.resourcemanager.batch.models.CheckNameAvailabilityParameters;
import com.azure.resourcemanager.batch.models.ResourceType;

/**
 * Samples for Location CheckNameAvailability.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-06-01/LocationCheckNameAvailability_AlreadyExists.json
     */
    /**
     * Sample code: LocationCheckNameAvailability_AlreadyExists.
     * 
     * @param manager Entry point to BatchManager.
     */
    public static void
        locationCheckNameAvailabilityAlreadyExists(com.azure.resourcemanager.batch.BatchManager manager) {
        manager.locations()
            .checkNameAvailabilityWithResponse("japaneast", new CheckNameAvailabilityParameters()
                .withName("existingaccountname").withType(ResourceType.MICROSOFT_BATCH_BATCH_ACCOUNTS),
                com.azure.core.util.Context.NONE);
    }
}
