
import com.azure.resourcemanager.eventgrid.models.NetworkSecurityPerimeterResourceType;

/**
 * Samples for NetworkSecurityPerimeterConfigurations Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/eventgrid/resource-manager/Microsoft.EventGrid/preview/2024-06-01-preview/examples/
     * NetworkSecurityPerimeterConfigurations_Get.json
     */
    /**
     * Sample code: NetworkSecurityPerimeterConfigurations_Get.
     * 
     * @param manager Entry point to EventGridManager.
     */
    public static void
        networkSecurityPerimeterConfigurationsGet(com.azure.resourcemanager.eventgrid.EventGridManager manager) {
        manager.networkSecurityPerimeterConfigurations().getWithResponse("examplerg",
            NetworkSecurityPerimeterResourceType.TOPICS, "exampleResourceName",
            "8f6b6269-84f2-4d09-9e31-1127efcd1e40perimeter", "someAssociation", com.azure.core.util.Context.NONE);
    }
}
