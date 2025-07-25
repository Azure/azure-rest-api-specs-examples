
import com.azure.resourcemanager.eventgrid.models.NetworkSecurityPerimeterResourceType;

/**
 * Samples for NetworkSecurityPerimeterConfigurations Reconcile.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/eventgrid/resource-manager/Microsoft.EventGrid/preview/2025-04-01-preview/examples/
     * NetworkSecurityPerimeterConfigurations_Reconcile.json
     */
    /**
     * Sample code: NetworkSecurityPerimeterConfigurations_Reconcile.
     * 
     * @param manager Entry point to EventGridManager.
     */
    public static void
        networkSecurityPerimeterConfigurationsReconcile(com.azure.resourcemanager.eventgrid.EventGridManager manager) {
        manager.networkSecurityPerimeterConfigurations().reconcile("examplerg",
            NetworkSecurityPerimeterResourceType.TOPICS, "exampleResourceName",
            "8f6b6269-84f2-4d09-9e31-1127efcd1e40perimeter", "someAssociation", com.azure.core.util.Context.NONE);
    }
}
