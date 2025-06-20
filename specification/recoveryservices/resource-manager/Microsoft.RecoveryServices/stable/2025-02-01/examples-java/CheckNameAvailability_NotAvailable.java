
import com.azure.resourcemanager.recoveryservices.models.CheckNameAvailabilityParameters;

/**
 * Samples for RecoveryServices CheckNameAvailability.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/recoveryservices/resource-manager/Microsoft.RecoveryServices/stable/2025-02-01/examples/
     * CheckNameAvailability_NotAvailable.json
     */
    /**
     * Sample code: Availability status of Resource Name when resource with same name, type and subscription exists.
     * 
     * @param manager Entry point to RecoveryServicesManager.
     */
    public static void availabilityStatusOfResourceNameWhenResourceWithSameNameTypeAndSubscriptionExists(
        com.azure.resourcemanager.recoveryservices.RecoveryServicesManager manager) {
        manager.recoveryServices()
            .checkNameAvailabilityWithResponse(
                "resGroupBar", "westus", new CheckNameAvailabilityParameters()
                    .withType("Microsoft.RecoveryServices/Vaults").withName("swaggerExample2"),
                com.azure.core.util.Context.NONE);
    }
}
