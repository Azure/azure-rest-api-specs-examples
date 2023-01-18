/** Samples for DataProtectionOperations List. */
public final class Main {
    /*
     * x-ms-original-file: specification/dataprotection/resource-manager/Microsoft.DataProtection/preview/2022-11-01-preview/examples/Operations/List.json
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
