
/**
 * Samples for ArcSettings CreateIdentity.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/azurestackhci/resource-manager/Microsoft.AzureStackHCI/StackHCI/stable/2024-04-01/examples/
     * CreateArcIdentity.json
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
