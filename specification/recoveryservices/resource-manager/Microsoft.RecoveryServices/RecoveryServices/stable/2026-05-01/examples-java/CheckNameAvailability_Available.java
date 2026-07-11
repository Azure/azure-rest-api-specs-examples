
import com.azure.resourcemanager.recoveryservices.models.CheckNameAvailabilityParameters;

/**
 * Samples for RecoveryServices CheckNameAvailability.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-05-01/CheckNameAvailability_Available.json
     */
    /**
     * Sample code: Availability status of Resource Name when no resource with same name, type and subscription exists,
     * nor has been deleted within last 24 hours.
     * 
     * @param manager Entry point to RecoveryServicesManager.
     */
    public static void
        availabilityStatusOfResourceNameWhenNoResourceWithSameNameTypeAndSubscriptionExistsNorHasBeenDeletedWithinLast24Hours(
            com.azure.resourcemanager.recoveryservices.RecoveryServicesManager manager) {
        manager.recoveryServices()
            .checkNameAvailabilityWithResponse(
                "resGroupFoo", "westus", new CheckNameAvailabilityParameters()
                    .withType("Microsoft.RecoveryServices/Vaults").withName("swaggerExample"),
                com.azure.core.util.Context.NONE);
    }
}
