
/**
 * Samples for Locations List.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/security/resource-manager/Microsoft.Security/preview/2015-06-01-preview/examples/Locations/
     * GetLocations_example.json
     */
    /**
     * Sample code: Get security data locations.
     * 
     * @param manager Entry point to SecurityManager.
     */
    public static void getSecurityDataLocations(com.azure.resourcemanager.security.SecurityManager manager) {
        manager.locations().list(com.azure.core.util.Context.NONE);
    }
}
