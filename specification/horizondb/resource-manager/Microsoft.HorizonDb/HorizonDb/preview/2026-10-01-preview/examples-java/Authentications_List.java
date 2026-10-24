
/**
 * Samples for Authentications List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-10-01-preview/Authentications_List.json
     */
    /**
     * Sample code: List the service-created authentication methods.
     * 
     * @param manager Entry point to HorizonDbManager.
     */
    public static void
        listTheServiceCreatedAuthenticationMethods(com.azure.resourcemanager.horizondb.HorizonDbManager manager) {
        manager.authentications().list("example-resource-group", "example-cluster", com.azure.core.util.Context.NONE);
    }
}
