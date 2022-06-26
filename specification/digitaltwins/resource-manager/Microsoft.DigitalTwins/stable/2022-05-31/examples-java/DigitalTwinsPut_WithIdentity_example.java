import com.azure.resourcemanager.digitaltwins.models.DigitalTwinsIdentity;
import com.azure.resourcemanager.digitaltwins.models.DigitalTwinsIdentityType;

/** Samples for DigitalTwins CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/digitaltwins/resource-manager/Microsoft.DigitalTwins/stable/2022-05-31/examples/DigitalTwinsPut_WithIdentity_example.json
     */
    /**
     * Sample code: Put a DigitalTwinsInstance resource with identity.
     *
     * @param manager Entry point to AzureDigitalTwinsManager.
     */
    public static void putADigitalTwinsInstanceResourceWithIdentity(
        com.azure.resourcemanager.digitaltwins.AzureDigitalTwinsManager manager) {
        manager
            .digitalTwins()
            .define("myDigitalTwinsService")
            .withRegion("WestUS2")
            .withExistingResourceGroup("resRg")
            .withIdentity(new DigitalTwinsIdentity().withType(DigitalTwinsIdentityType.SYSTEM_ASSIGNED))
            .create();
    }
}
