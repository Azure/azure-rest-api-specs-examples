
import com.azure.resourcemanager.customerinsights.models.Hub;
import com.azure.resourcemanager.customerinsights.models.HubBillingInfoFormat;

/**
 * Samples for Hubs Update.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/customer-insights/resource-manager/Microsoft.CustomerInsights/stable/2017-04-26/examples/HubsUpdate
     * .json
     */
    /**
     * Sample code: Hubs_Update.
     * 
     * @param manager Entry point to CustomerInsightsManager.
     */
    public static void hubsUpdate(com.azure.resourcemanager.customerinsights.CustomerInsightsManager manager) {
        Hub resource = manager.hubs()
            .getByResourceGroupWithResponse("TestHubRG", "sdkTestHub", com.azure.core.util.Context.NONE).getValue();
        resource.update()
            .withHubBillingInfo(new HubBillingInfoFormat().withSkuName("B0").withMinUnits(1).withMaxUnits(5)).apply();
    }
}
