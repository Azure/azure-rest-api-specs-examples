
import com.azure.resourcemanager.sql.fluent.models.MaintenanceWindowsInner;
import com.azure.resourcemanager.sql.models.DayOfWeek;
import com.azure.resourcemanager.sql.models.MaintenanceWindowTimeRange;
import java.util.Arrays;

/**
 * Samples for MaintenanceWindowsOperation CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/sql/resource-manager/Microsoft.Sql/stable/2021-11-01/examples/CreateOrUpdateMaintenanceWindows.json
     */
    /**
     * Sample code: Sets maintenance window settings for a selected database.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void
        setsMaintenanceWindowSettingsForASelectedDatabase(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.sqlServers().manager().serviceClient().getMaintenanceWindowsOperations().createOrUpdateWithResponse(
            "Default-SQL-SouthEastAsia", "testsvr", "testdwdb", "current",
            new MaintenanceWindowsInner().withTimeRanges(Arrays.asList(new MaintenanceWindowTimeRange()
                .withDayOfWeek(DayOfWeek.SATURDAY).withStartTime("00:00:00").withDuration("PT60M"))),
            com.azure.core.util.Context.NONE);
    }
}
