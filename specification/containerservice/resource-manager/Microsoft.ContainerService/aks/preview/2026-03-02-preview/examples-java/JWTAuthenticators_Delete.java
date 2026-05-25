
/**
 * Samples for JWTAuthenticators Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-02-preview/JWTAuthenticators_Delete.json
     */
    /**
     * Sample code: Delete JWT Authenticator.
     * 
     * @param manager Entry point to ContainerServiceManager.
     */
    public static void
        deleteJWTAuthenticator(com.azure.resourcemanager.containerservice.ContainerServiceManager manager) {
        manager.serviceClient().getJWTAuthenticators().delete("rg1", "clustername1", "jwt1",
            com.azure.core.util.Context.NONE);
    }
}
