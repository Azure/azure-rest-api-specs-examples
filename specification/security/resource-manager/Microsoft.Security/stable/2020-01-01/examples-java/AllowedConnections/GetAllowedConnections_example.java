
import com.azure.resourcemanager.security.models.ConnectionType;

/**
 * Samples for AllowedConnections Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/security/resource-manager/Microsoft.Security/stable/2020-01-01/examples/AllowedConnections/
     * GetAllowedConnections_example.json
     */
    /**
     * Sample code: Get allowed connections.
     * 
     * @param manager Entry point to SecurityManager.
     */
    public static void getAllowedConnections(com.azure.resourcemanager.security.SecurityManager manager) {
        manager.allowedConnections().getWithResponse("myResourceGroup", "centralus", ConnectionType.INTERNAL,
            com.azure.core.util.Context.NONE);
    }
}
