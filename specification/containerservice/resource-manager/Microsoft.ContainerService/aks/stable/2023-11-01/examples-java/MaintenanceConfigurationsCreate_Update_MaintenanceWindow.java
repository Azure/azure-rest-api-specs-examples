
import com.azure.resourcemanager.containerservice.fluent.models.MaintenanceConfigurationInner;
import com.azure.resourcemanager.containerservice.models.DateSpan;
import com.azure.resourcemanager.containerservice.models.MaintenanceWindow;
import com.azure.resourcemanager.containerservice.models.RelativeMonthlySchedule;
import com.azure.resourcemanager.containerservice.models.Schedule;
import com.azure.resourcemanager.containerservice.models.Type;
import com.azure.resourcemanager.containerservice.models.WeekDay;
import java.time.LocalDate;
import java.util.Arrays;

/**
 * Samples for MaintenanceConfigurations CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/containerservice/resource-manager/Microsoft.ContainerService/aks/stable/2023-11-01/examples/
     * MaintenanceConfigurationsCreate_Update_MaintenanceWindow.json
     */
    /**
     * Sample code: Create/Update Maintenance Configuration with Maintenance Window.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void createUpdateMaintenanceConfigurationWithMaintenanceWindow(
        com.azure.resourcemanager.AzureResourceManager azure) {
        azure.kubernetesClusters().manager().serviceClient().getMaintenanceConfigurations().createOrUpdateWithResponse(
            "rg1", "clustername1", "aksManagedAutoUpgradeSchedule",
            new MaintenanceConfigurationInner().withMaintenanceWindow(new MaintenanceWindow()
                .withSchedule(new Schedule().withRelativeMonthly(new RelativeMonthlySchedule().withIntervalMonths(3)
                    .withWeekIndex(Type.FIRST).withDayOfWeek(WeekDay.MONDAY)))
                .withDurationHours(10).withUtcOffset("+05:30").withStartDate(LocalDate.parse("2023-01-01"))
                .withStartTime("08:30")
                .withNotAllowedDates(Arrays.asList(
                    new DateSpan().withStart(LocalDate.parse("2023-02-18")).withEnd(LocalDate.parse("2023-02-25")),
                    new DateSpan().withStart(LocalDate.parse("2023-12-23")).withEnd(LocalDate.parse("2024-01-05"))))),
            com.azure.core.util.Context.NONE);
    }
}
