
/**
 * Samples for Locations Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/security/resource-manager/Microsoft.Security/preview/2015-06-01-preview/examples/Locations/
     * GetLocation_example.json
     */
    /**
     * Sample code: Get security data location.
     * 
     * @param manager Entry point to SecurityManager.
     */
    public static void getSecurityDataLocation(com.azure.resourcemanager.security.SecurityManager manager) {
        manager.locations().getWithResponse("centralus", com.azure.core.util.Context.NONE);
    }
}
