import com.azure.core.util.Context;
import com.azure.resourcemanager.batch.models.CheckNameAvailabilityParameters;

/** Samples for Location CheckNameAvailability. */
public final class Main {
    /*
     * x-ms-original-file: specification/batch/resource-manager/Microsoft.Batch/stable/2022-10-01/examples/LocationCheckNameAvailability_Available.json
     */
    /**
     * Sample code: LocationCheckNameAvailability_Available.
     *
     * @param manager Entry point to BatchManager.
     */
    public static void locationCheckNameAvailabilityAvailable(com.azure.resourcemanager.batch.BatchManager manager) {
        manager
            .locations()
            .checkNameAvailabilityWithResponse(
                "japaneast", new CheckNameAvailabilityParameters().withName("newaccountname"), Context.NONE);
    }
}
