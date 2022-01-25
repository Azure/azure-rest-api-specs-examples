Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-synapse_1.0.0-beta.5/sdk/synapse/azure-resourcemanager-synapse/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;

/** Samples for SqlPoolMaintenanceWindows Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/synapse/resource-manager/Microsoft.Synapse/stable/2021-06-01/examples/GetMaintenanceWindows.json
     */
    /**
     * Sample code: Gets maintenance window settings for a selected SQL Analytics pool.
     *
     * @param manager Entry point to SynapseManager.
     */
    public static void getsMaintenanceWindowSettingsForASelectedSQLAnalyticsPool(
        com.azure.resourcemanager.synapse.SynapseManager manager) {
        manager
            .sqlPoolMaintenanceWindows()
            .getWithResponse("samplerg", "testworkspace", "testsp", "current", Context.NONE);
    }
}
```
