```java
import com.azure.core.util.Context;
import com.azure.resourcemanager.synapse.fluent.models.MaintenanceWindowsInner;
import com.azure.resourcemanager.synapse.models.DayOfWeek;
import com.azure.resourcemanager.synapse.models.MaintenanceWindowTimeRange;
import java.util.Arrays;

/** Samples for SqlPoolMaintenanceWindows CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/synapse/resource-manager/Microsoft.Synapse/stable/2021-06-01/examples/CreateOrUpdateMaintenanceWindows.json
     */
    /**
     * Sample code: Sets maintenance window settings for a selected SQL Analytics Pool.
     *
     * @param manager Entry point to SynapseManager.
     */
    public static void setsMaintenanceWindowSettingsForASelectedSQLAnalyticsPool(
        com.azure.resourcemanager.synapse.SynapseManager manager) {
        manager
            .sqlPoolMaintenanceWindows()
            .createOrUpdateWithResponse(
                "samplerg",
                "testworkspace",
                "testsp",
                "current",
                new MaintenanceWindowsInner()
                    .withTimeRanges(
                        Arrays
                            .asList(
                                new MaintenanceWindowTimeRange()
                                    .withDayOfWeek(DayOfWeek.SATURDAY)
                                    .withStartTime("00:00:00")
                                    .withDuration("PT60M"))),
                Context.NONE);
    }
}
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-synapse_1.0.0-beta.6/sdk/synapse/azure-resourcemanager-synapse/README.md) on how to add the SDK to your project and authenticate.
