
/**
 * Samples for JWTAuthenticators Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-02-preview/JWTAuthenticators_Get.json
     */
    /**
     * Sample code: Get JWT Authenticator.
     * 
     * @param manager Entry point to ContainerServiceManager.
     */
    public static void getJWTAuthenticator(com.azure.resourcemanager.containerservice.ContainerServiceManager manager) {
        manager.serviceClient().getJWTAuthenticators().getWithResponse("rg1", "clustername1", "jwt1",
            com.azure.core.util.Context.NONE);
    }
}
