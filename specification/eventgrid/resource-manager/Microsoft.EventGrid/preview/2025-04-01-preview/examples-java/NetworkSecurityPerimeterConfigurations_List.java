
import com.azure.resourcemanager.eventgrid.models.NetworkSecurityPerimeterResourceType;

/**
 * Samples for NetworkSecurityPerimeterConfigurations List.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/eventgrid/resource-manager/Microsoft.EventGrid/preview/2025-04-01-preview/examples/
     * NetworkSecurityPerimeterConfigurations_List.json
     */
    /**
     * Sample code: NetworkSecurityPerimeterConfigurations_List.
     * 
     * @param manager Entry point to EventGridManager.
     */
    public static void
        networkSecurityPerimeterConfigurationsList(com.azure.resourcemanager.eventgrid.EventGridManager manager) {
        manager.networkSecurityPerimeterConfigurations().list("examplerg", NetworkSecurityPerimeterResourceType.TOPICS,
            "exampleResourceName", com.azure.core.util.Context.NONE);
    }
}
