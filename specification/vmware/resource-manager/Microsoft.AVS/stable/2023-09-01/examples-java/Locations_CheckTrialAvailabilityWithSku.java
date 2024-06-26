
import com.azure.resourcemanager.avs.models.Sku;

/**
 * Samples for Locations CheckTrialAvailability.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/vmware/resource-manager/Microsoft.AVS/stable/2023-09-01/examples/
     * Locations_CheckTrialAvailabilityWithSku.json
     */
    /**
     * Sample code: Locations_CheckTrialAvailabilityWithSku.
     * 
     * @param manager Entry point to AvsManager.
     */
    public static void locationsCheckTrialAvailabilityWithSku(com.azure.resourcemanager.avs.AvsManager manager) {
        manager.locations().checkTrialAvailabilityWithResponse("eastus", new Sku().withName("avs52t"),
            com.azure.core.util.Context.NONE);
    }
}
