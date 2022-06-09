```java
import com.azure.core.util.Context;

/** Samples for Grafana List. */
public final class Main {
    /*
     * x-ms-original-file: specification/dashboard/resource-manager/Microsoft.Dashboard/preview/2021-09-01-preview/examples/Grafana_List.json
     */
    /**
     * Sample code: Grafana_List.
     *
     * @param manager Entry point to DashboardManager.
     */
    public static void grafanaList(com.azure.resourcemanager.dashboard.DashboardManager manager) {
        manager.grafanas().list(Context.NONE);
    }
}
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-dashboard_1.0.0-beta.1/sdk/dashboard/azure-resourcemanager-dashboard/README.md) on how to add the SDK to your project and authenticate.
