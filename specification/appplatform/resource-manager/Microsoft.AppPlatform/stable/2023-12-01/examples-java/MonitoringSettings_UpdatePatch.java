
import com.azure.resourcemanager.appplatform.fluent.models.MonitoringSettingResourceInner;
import com.azure.resourcemanager.appplatform.models.MonitoringSettingProperties;

/**
 * Samples for MonitoringSettings UpdatePatch.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/appplatform/resource-manager/Microsoft.AppPlatform/stable/2023-12-01/examples/
     * MonitoringSettings_UpdatePatch.json
     */
    /**
     * Sample code: MonitoringSettings_UpdatePatch.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void monitoringSettingsUpdatePatch(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.springServices().manager().serviceClient().getMonitoringSettings().updatePatch("myResourceGroup",
            "myservice",
            new MonitoringSettingResourceInner().withProperties(new MonitoringSettingProperties().withTraceEnabled(true)
                .withAppInsightsInstrumentationKey("fakeTokenPlaceholder").withAppInsightsSamplingRate(10.0D)),
            com.azure.core.util.Context.NONE);
    }
}
