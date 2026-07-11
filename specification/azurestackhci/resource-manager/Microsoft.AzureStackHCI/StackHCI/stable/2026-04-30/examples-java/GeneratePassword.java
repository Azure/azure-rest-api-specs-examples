
/**
 * Samples for ArcSettings GeneratePassword.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-04-30/GeneratePassword.json
     */
    /**
     * Sample code: Generate Password.
     * 
     * @param manager Entry point to AzureStackHciManager.
     */
    public static void generatePassword(com.azure.resourcemanager.azurestackhci.AzureStackHciManager manager) {
        manager.arcSettings().generatePasswordWithResponse("test-rg", "myCluster", "default",
            com.azure.core.util.Context.NONE);
    }
}
