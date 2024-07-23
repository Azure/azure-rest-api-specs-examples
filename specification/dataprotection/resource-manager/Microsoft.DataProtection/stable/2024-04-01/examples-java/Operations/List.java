
/**
 * Samples for DataProtectionOperations List.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/dataprotection/resource-manager/Microsoft.DataProtection/stable/2024-04-01/examples/Operations/List
     * .json
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
