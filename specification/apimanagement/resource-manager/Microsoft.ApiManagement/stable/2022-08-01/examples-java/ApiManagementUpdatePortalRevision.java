
import com.azure.resourcemanager.apimanagement.models.PortalRevisionContract;

/**
 * Samples for PortalRevision Update.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2022-08-01/examples/
     * ApiManagementUpdatePortalRevision.json
     */
    /**
     * Sample code: ApiManagementUpdatePortalRevision.
     * 
     * @param manager Entry point to ApiManagementManager.
     */
    public static void
        apiManagementUpdatePortalRevision(com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        PortalRevisionContract resource = manager.portalRevisions()
            .getWithResponse("rg1", "apimService1", "20201112101010", com.azure.core.util.Context.NONE).getValue();
        resource.update().withDescription("portal revision update").withIsCurrent(true).withIfMatch("*").apply();
    }
}
