/** Samples for DigitalTwins ListByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/digitaltwins/resource-manager/Microsoft.DigitalTwins/stable/2023-01-31/examples/DigitalTwinsListByResourceGroup_example.json
     */
    /**
     * Sample code: Get DigitalTwinsInstance resources by resource group.
     *
     * @param manager Entry point to AzureDigitalTwinsManager.
     */
    public static void getDigitalTwinsInstanceResourcesByResourceGroup(
        com.azure.resourcemanager.digitaltwins.AzureDigitalTwinsManager manager) {
        manager.digitalTwins().listByResourceGroup("resRg", com.azure.core.util.Context.NONE);
    }
}
