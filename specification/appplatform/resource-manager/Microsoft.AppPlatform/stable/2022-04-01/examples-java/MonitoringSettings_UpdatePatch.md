Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager_2.15.0/sdk/resourcemanager/azure-resourcemanager/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;
import com.azure.resourcemanager.appplatform.fluent.models.MonitoringSettingResourceInner;
import com.azure.resourcemanager.appplatform.models.MonitoringSettingProperties;

/** Samples for MonitoringSettings UpdatePatch. */
public final class Main {
    /*
     * x-ms-original-file: specification/appplatform/resource-manager/Microsoft.AppPlatform/stable/2022-04-01/examples/MonitoringSettings_UpdatePatch.json
     */
    /**
     * Sample code: MonitoringSettings_UpdatePatch.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void monitoringSettingsUpdatePatch(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .springServices()
            .manager()
            .serviceClient()
            .getMonitoringSettings()
            .updatePatch(
                "myResourceGroup",
                "myservice",
                new MonitoringSettingResourceInner()
                    .withProperties(
                        new MonitoringSettingProperties()
                            .withTraceEnabled(true)
                            .withAppInsightsInstrumentationKey("00000000-0000-0000-0000-000000000000")
                            .withAppInsightsSamplingRate(10.0)),
                Context.NONE);
    }
}
```
