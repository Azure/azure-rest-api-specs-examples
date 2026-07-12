
/**
 * Samples for ArcSettings CreateIdentity.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-04-30/CreateArcIdentity.json
     */
    /**
     * Sample code: Create Arc Identity.
     * 
     * @param manager Entry point to AzureStackHciManager.
     */
    public static void createArcIdentity(com.azure.resourcemanager.azurestackhci.AzureStackHciManager manager) {
        manager.arcSettings().createIdentity("test-rg", "myCluster", "default", com.azure.core.util.Context.NONE);
    }
}
