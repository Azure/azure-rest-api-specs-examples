
import com.azure.resourcemanager.appcontainers.models.CheckMigrationEligibilityRequest;

/**
 * Samples for ManagedEnvironments CheckMigrationEligibility.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-07-01/ManagedEnvironments_CheckMigrationEligibility_Ineligible.json
     */
    /**
     * Sample code: Check migration eligibility for an ineligible managed environment.
     * 
     * @param manager Entry point to ContainerAppsApiManager.
     */
    public static void checkMigrationEligibilityForAnIneligibleManagedEnvironment(
        com.azure.resourcemanager.appcontainers.ContainerAppsApiManager manager) {
        manager.managedEnvironments().checkMigrationEligibilityWithResponse("examplerg", "my-environment",
            new CheckMigrationEligibilityRequest().withTargetMode("Express"), com.azure.core.util.Context.NONE);
    }
}
