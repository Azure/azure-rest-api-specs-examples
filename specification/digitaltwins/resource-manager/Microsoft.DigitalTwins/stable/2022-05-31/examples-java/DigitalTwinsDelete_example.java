import com.azure.core.util.Context;

/** Samples for DigitalTwins Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/digitaltwins/resource-manager/Microsoft.DigitalTwins/stable/2022-05-31/examples/DigitalTwinsDelete_example.json
     */
    /**
     * Sample code: Delete a DigitalTwinsInstance resource.
     *
     * @param manager Entry point to AzureDigitalTwinsManager.
     */
    public static void deleteADigitalTwinsInstanceResource(
        com.azure.resourcemanager.digitaltwins.AzureDigitalTwinsManager manager) {
        manager.digitalTwins().delete("resRg", "myDigitalTwinsService", Context.NONE);
    }
}
