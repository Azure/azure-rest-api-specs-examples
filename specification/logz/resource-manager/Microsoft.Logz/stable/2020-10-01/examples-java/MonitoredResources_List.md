Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-logz_1.0.0-beta.1/sdk/logz/azure-resourcemanager-logz/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;

/** Samples for Monitors ListMonitoredResources. */
public final class Main {
    /*
     * x-ms-original-file: specification/logz/resource-manager/Microsoft.Logz/stable/2020-10-01/examples/MonitoredResources_List.json
     */
    /**
     * Sample code: MonitoredResources_List.
     *
     * @param manager Entry point to LogzManager.
     */
    public static void monitoredResourcesList(com.azure.resourcemanager.logz.LogzManager manager) {
        manager.monitors().listMonitoredResources("myResourceGroup", "myMonitor", Context.NONE);
    }
}
```
