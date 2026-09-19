
import com.azure.resourcemanager.appcontainers.models.ScheduledEntry;
import com.azure.resourcemanager.appcontainers.models.WeekDay;
import java.util.Arrays;

/**
 * Samples for MaintenanceConfigurations CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-07-01/ManagedEnvironment_MaintenanceConfigurations_CreateOrUpdate.json
     */
    /**
     * Sample code: ManagedEnvironmentMaintenanceConfigurationsCreateOrUpdate.
     * 
     * @param manager Entry point to ContainerAppsApiManager.
     */
    public static void managedEnvironmentMaintenanceConfigurationsCreateOrUpdate(
        com.azure.resourcemanager.appcontainers.ContainerAppsApiManager manager) {
        manager.maintenanceConfigurations().define("default").withExistingManagedEnvironment("rg1", "managedEnv")
            .withScheduledEntries(Arrays
                .asList(new ScheduledEntry().withWeekDay(WeekDay.SUNDAY).withStartHourUtc(12).withDurationHours(9)))
            .create();
    }
}
