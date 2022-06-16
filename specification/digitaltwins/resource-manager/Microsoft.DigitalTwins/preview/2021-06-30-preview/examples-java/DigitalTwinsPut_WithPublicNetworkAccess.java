import com.azure.resourcemanager.digitaltwins.models.PublicNetworkAccess;

/** Samples for DigitalTwins CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/digitaltwins/resource-manager/Microsoft.DigitalTwins/preview/2021-06-30-preview/examples/DigitalTwinsPut_WithPublicNetworkAccess.json
     */
    /**
     * Sample code: Put a DigitalTwinsInstance resource with publicNetworkAccess property.
     *
     * @param manager Entry point to AzureDigitalTwinsManager.
     */
    public static void putADigitalTwinsInstanceResourceWithPublicNetworkAccessProperty(
        com.azure.resourcemanager.digitaltwins.AzureDigitalTwinsManager manager) {
        manager
            .digitalTwins()
            .define("myDigitalTwinsService")
            .withRegion("WestUS2")
            .withExistingResourceGroup("resRg")
            .withPublicNetworkAccess(PublicNetworkAccess.ENABLED)
            .create();
    }
}
