
/**
 * Samples for Operations List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-08-01/Operations_List.json
     */
    /**
     * Sample code: ListOperations.
     * 
     * @param manager Entry point to ServiceGroupsManager.
     */
    public static void listOperations(com.azure.resourcemanager.servicegroups.ServiceGroupsManager manager) {
        manager.operations().list(com.azure.core.util.Context.NONE);
    }
}
