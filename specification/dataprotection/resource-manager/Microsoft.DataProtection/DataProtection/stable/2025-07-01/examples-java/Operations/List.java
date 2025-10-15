
/**
 * Samples for DataProtectionOperations List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/Operations/List.json
     */
    /**
     * Sample code: Returns the list of supported REST operations.
     * 
     * @param manager Entry point to DataProtectionManager.
     */
    public static void returnsTheListOfSupportedRESTOperations(
        com.azure.resourcemanager.dataprotection.DataProtectionManager manager) {
        manager.dataProtectionOperations().list(com.azure.core.util.Context.NONE);
    }
}
