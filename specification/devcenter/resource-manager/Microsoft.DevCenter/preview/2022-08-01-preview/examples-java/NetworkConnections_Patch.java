import com.azure.core.util.Context;
import com.azure.resourcemanager.devcenter.models.NetworkConnection;

/** Samples for NetworkConnections Update. */
public final class Main {
    /*
     * x-ms-original-file: specification/devcenter/resource-manager/Microsoft.DevCenter/preview/2022-08-01-preview/examples/NetworkConnections_Patch.json
     */
    /**
     * Sample code: NetworkConnections_Update.
     *
     * @param manager Entry point to DevCenterManager.
     */
    public static void networkConnectionsUpdate(com.azure.resourcemanager.devcenter.DevCenterManager manager) {
        NetworkConnection resource =
            manager
                .networkConnections()
                .getByResourceGroupWithResponse("rg1", "uswest3network", Context.NONE)
                .getValue();
        resource.update().withDomainPassword("New Password value for user").apply();
    }
}
