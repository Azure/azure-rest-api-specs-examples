
import com.azure.resourcemanager.containerservice.fluent.models.MaintenanceConfigurationInner;
import com.azure.resourcemanager.containerservice.models.TimeInWeek;
import com.azure.resourcemanager.containerservice.models.TimeSpan;
import com.azure.resourcemanager.containerservice.models.WeekDay;
import java.time.OffsetDateTime;
import java.util.Arrays;

/**
 * Samples for MaintenanceConfigurations CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/containerservice/resource-manager/Microsoft.ContainerService/aks/stable/2025-04-01/examples/
     * MaintenanceConfigurationsCreate_Update.json
     */
    /**
     * Sample code: Create/Update Maintenance Configuration.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void createUpdateMaintenanceConfiguration(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.kubernetesClusters().manager().serviceClient().getMaintenanceConfigurations()
            .createOrUpdateWithResponse("rg1", "clustername1", "default", new MaintenanceConfigurationInner()
                .withTimeInWeek(
                    Arrays.asList(new TimeInWeek().withDay(WeekDay.MONDAY).withHourSlots(Arrays.asList(1, 2))))
                .withNotAllowedTime(Arrays.asList(new TimeSpan().withStart(OffsetDateTime.parse("2020-11-26T03:00:00Z"))
                    .withEnd(OffsetDateTime.parse("2020-11-30T12:00:00Z")))),
                com.azure.core.util.Context.NONE);
    }
}
