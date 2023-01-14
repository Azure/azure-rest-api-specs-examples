import com.azure.resourcemanager.customerinsights.models.HubBillingInfoFormat;

/** Samples for Hubs CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/customer-insights/resource-manager/Microsoft.CustomerInsights/stable/2017-04-26/examples/HubsCreateOrUpdate.json
     */
    /**
     * Sample code: Hubs_CreateOrUpdate.
     *
     * @param manager Entry point to CustomerInsightsManager.
     */
    public static void hubsCreateOrUpdate(com.azure.resourcemanager.customerinsights.CustomerInsightsManager manager) {
        manager
            .hubs()
            .define("sdkTestHub")
            .withRegion("West US")
            .withExistingResourceGroup("TestHubRG")
            .withHubBillingInfo(new HubBillingInfoFormat().withSkuName("B0").withMinUnits(1).withMaxUnits(5))
            .create();
    }
}
