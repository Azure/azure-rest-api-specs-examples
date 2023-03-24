/** Samples for DigitalTwins GetByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/digitaltwins/resource-manager/Microsoft.DigitalTwins/stable/2023-01-31/examples/DigitalTwinsGet_example.json
     */
    /**
     * Sample code: Get a DigitalTwinsInstance resource.
     *
     * @param manager Entry point to AzureDigitalTwinsManager.
     */
    public static void getADigitalTwinsInstanceResource(
        com.azure.resourcemanager.digitaltwins.AzureDigitalTwinsManager manager) {
        manager
            .digitalTwins()
            .getByResourceGroupWithResponse("resRg", "myDigitalTwinsService", com.azure.core.util.Context.NONE);
    }
}
