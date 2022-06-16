import com.azure.core.util.Context;
import com.azure.resourcemanager.appplatform.fluent.models.MonitoringSettingResourceInner;
import com.azure.resourcemanager.appplatform.models.MonitoringSettingProperties;

/** Samples for MonitoringSettings UpdatePut. */
public final class Main {
    /*
     * x-ms-original-file: specification/appplatform/resource-manager/Microsoft.AppPlatform/stable/2022-04-01/examples/MonitoringSettings_UpdatePut.json
     */
    /**
     * Sample code: MonitoringSettings_UpdatePut.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void monitoringSettingsUpdatePut(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .springServices()
            .manager()
            .serviceClient()
            .getMonitoringSettings()
            .updatePut(
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
