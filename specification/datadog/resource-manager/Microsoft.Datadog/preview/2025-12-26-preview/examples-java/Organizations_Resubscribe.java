
import com.azure.resourcemanager.datadog.models.ResourceSku;
import com.azure.resourcemanager.datadog.models.ResubscribeProperties;

/**
 * Samples for Organizations Resubscribe.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-12-26-preview/Organizations_Resubscribe.json
     */
    /**
     * Sample code: Organizations_Resubscribe.
     * 
     * @param manager Entry point to MicrosoftDatadogManager.
     */
    public static void organizationsResubscribe(com.azure.resourcemanager.datadog.MicrosoftDatadogManager manager) {
        manager.organizations().resubscribe("myResourceGroup", "myMonitor",
            new ResubscribeProperties().withSku(new ResourceSku().withName("planName"))
                .withAzureSubscriptionId("subscriptionId").withResourceGroup("resourceGroup"),
            com.azure.core.util.Context.NONE);
    }
}
