
/**
 * Samples for Operations List.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/security/resource-manager/Microsoft.Security/preview/2024-10-01-preview/examples/Operations/
     * ListOperations_example.json
     */
    /**
     * Sample code: List the operations for the Microsoft.Security (Microsoft Defender for Cloud) resource provider.
     * 
     * @param manager Entry point to SecurityManager.
     */
    public static void listTheOperationsForTheMicrosoftSecurityMicrosoftDefenderForCloudResourceProvider(
        com.azure.resourcemanager.security.SecurityManager manager) {
        manager.operations().list(com.azure.core.util.Context.NONE);
    }
}
