
import com.azure.resourcemanager.labservices.models.EnableState;
import com.azure.resourcemanager.labservices.models.Lab;
import com.azure.resourcemanager.labservices.models.SecurityProfile;

/**
 * Samples for Labs Update.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/labservices/resource-manager/Microsoft.LabServices/stable/2022-08-01/examples/Labs/patchLab.json
     */
    /**
     * Sample code: patchLab.
     * 
     * @param manager Entry point to LabServicesManager.
     */
    public static void patchLab(com.azure.resourcemanager.labservices.LabServicesManager manager) {
        Lab resource = manager.labs()
            .getByResourceGroupWithResponse("testrg123", "testlab", com.azure.core.util.Context.NONE).getValue();
        resource.update().withSecurityProfile(new SecurityProfile().withOpenAccess(EnableState.ENABLED)).apply();
    }
}
