
import com.azure.resourcemanager.labservices.models.ConnectionProfile;
import com.azure.resourcemanager.labservices.models.ConnectionType;
import com.azure.resourcemanager.labservices.models.LabPlan;

/**
 * Samples for LabPlans Update.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/labservices/resource-manager/Microsoft.LabServices/stable/2022-08-01/examples/LabPlans/patchLabPlan
     * .json
     */
    /**
     * Sample code: patchLabPlan.
     * 
     * @param manager Entry point to LabServicesManager.
     */
    public static void patchLabPlan(com.azure.resourcemanager.labservices.LabServicesManager manager) {
        LabPlan resource = manager.labPlans()
            .getByResourceGroupWithResponse("testrg123", "testlabplan", com.azure.core.util.Context.NONE).getValue();
        resource.update()
            .withDefaultConnectionProfile(
                new ConnectionProfile().withWebSshAccess(ConnectionType.NONE).withWebRdpAccess(ConnectionType.NONE)
                    .withClientSshAccess(ConnectionType.PUBLIC).withClientRdpAccess(ConnectionType.PUBLIC))
            .apply();
    }
}
