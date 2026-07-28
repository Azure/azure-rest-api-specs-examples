
import com.azure.resourcemanager.datadog.models.SaaSData;

/**
 * Samples for DatadogMonitorResources LinkSaaS.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-12-26-preview/Monitors_LinkSaaS.json
     */
    /**
     * Sample code: Monitors_LinkSaaS.
     * 
     * @param manager Entry point to MicrosoftDatadogManager.
     */
    public static void monitorsLinkSaaS(com.azure.resourcemanager.datadog.MicrosoftDatadogManager manager) {
        manager.datadogMonitorResources().linkSaaS("myResourceGroup", "myMonitor", new SaaSData().withSaaSResourceId(
            "/subscriptions/1a2b3c4d-5e6f-7a8b-9c0d-e1f2a3b4c5d6/resourceGroups/myResourceGroup/providers/Microsoft.SaaS/resources/mySaaSResource"),
            com.azure.core.util.Context.NONE);
    }
}
