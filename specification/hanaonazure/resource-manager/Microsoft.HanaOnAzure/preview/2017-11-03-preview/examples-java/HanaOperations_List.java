/** Samples for Operations List. */
public final class Main {
    /*
     * x-ms-original-file: specification/hanaonazure/resource-manager/Microsoft.HanaOnAzure/preview/2017-11-03-preview/examples/HanaOperations_List.json
     */
    /**
     * Sample code: List all HANA management operations supported by HANA RP.
     *
     * @param manager Entry point to HanaManager.
     */
    public static void listAllHANAManagementOperationsSupportedByHANARP(
        com.azure.resourcemanager.hanaonazure.HanaManager manager) {
        manager.operations().list(com.azure.core.util.Context.NONE);
    }
}
