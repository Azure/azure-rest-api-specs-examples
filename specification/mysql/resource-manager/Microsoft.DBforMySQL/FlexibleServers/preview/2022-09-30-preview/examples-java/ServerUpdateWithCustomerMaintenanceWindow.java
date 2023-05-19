import com.azure.resourcemanager.mysqlflexibleserver.models.MaintenanceWindow;
import com.azure.resourcemanager.mysqlflexibleserver.models.Server;
import java.util.HashMap;
import java.util.Map;

/** Samples for Servers Update. */
public final class Main {
    /*
     * x-ms-original-file: specification/mysql/resource-manager/Microsoft.DBforMySQL/FlexibleServers/preview/2022-09-30-preview/examples/ServerUpdateWithCustomerMaintenanceWindow.json
     */
    /**
     * Sample code: Update server customer maintenance window.
     *
     * @param manager Entry point to MySqlManager.
     */
    public static void updateServerCustomerMaintenanceWindow(
        com.azure.resourcemanager.mysqlflexibleserver.MySqlManager manager) {
        Server resource =
            manager
                .servers()
                .getByResourceGroupWithResponse("testrg", "mysqltestserver", com.azure.core.util.Context.NONE)
                .getValue();
        resource
            .update()
            .withMaintenanceWindow(
                new MaintenanceWindow()
                    .withCustomWindow("Enabled")
                    .withStartHour(8)
                    .withStartMinute(0)
                    .withDayOfWeek(1))
            .apply();
    }

    @SuppressWarnings("unchecked")
    private static <T> Map<String, T> mapOf(Object... inputs) {
        Map<String, T> map = new HashMap<>();
        for (int i = 0; i < inputs.length; i += 2) {
            String key = (String) inputs[i];
            T value = (T) inputs[i + 1];
            map.put(key, value);
        }
        return map;
    }
}
