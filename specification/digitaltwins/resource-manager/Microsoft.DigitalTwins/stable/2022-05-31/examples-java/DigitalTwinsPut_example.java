/** Samples for DigitalTwins CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/digitaltwins/resource-manager/Microsoft.DigitalTwins/stable/2022-05-31/examples/DigitalTwinsPut_example.json
     */
    /**
     * Sample code: Put a DigitalTwinsInstance resource.
     *
     * @param manager Entry point to AzureDigitalTwinsManager.
     */
    public static void putADigitalTwinsInstanceResource(
        com.azure.resourcemanager.digitaltwins.AzureDigitalTwinsManager manager) {
        manager
            .digitalTwins()
            .define("myDigitalTwinsService")
            .withRegion("WestUS2")
            .withExistingResourceGroup("resRg")
            .create();
    }
}
